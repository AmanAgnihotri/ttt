// Copyright 2024 Aman Agnihotri
// SPDX-License-Identifier: AGPL-3.0-only

package web

import (
	"sync"

	"ttt/internal/base/domain/game"
	"ttt/internal/base/domain/game/user"
	"ttt/internal/game/common/domain/board"
	"ttt/internal/game/play/create/api"
	"ttt/internal/game/play/create/domain"
	"ttt/pkg/infra/web/time"
)

const maxUserCount = 2

type Controller struct {
	timer       time.Timer
	store       domain.Store
	gameID      game.ID
	done        func()
	connections map[user.ID]*Conn
	handler     api.Handler
	running     bool
	sync.RWMutex
}

func NewController(
	timer time.Timer,
	store domain.Store,
	gameID game.ID,
	done func(),
) *Controller {
	return &Controller{
		timer:       timer,
		store:       store,
		gameID:      gameID,
		done:        done,
		connections: make(map[user.ID]*Conn, maxUserCount),
		handler:     nil,
		running:     false,
		RWMutex:     sync.RWMutex{},
	}
}

func (c *Controller) Handle(userID user.ID, conn *Conn) {
	c.create(userID, conn)

	conn.OnClose(func(ctrl *Controller, userID user.ID) {
		ctrl.delete(userID)
		ctrl.tryCleanup()
	})

	conn.OnRequest(func(conn *Conn, ctrl *Controller, req request) {
		ctrl.RLock()
		handler := ctrl.handler
		ctrl.RUnlock()

		if handler == nil {
			conn.CloseAsInvalidMessage()
		}

		if command, ok := ctrl.newCommand(req); ok {
			handler.Handle(command)
		} else {
			conn.CloseAsInvalidMessage()
		}
	})

	c.RLock()
	running := c.running
	c.RUnlock()

	if running {
		conn.WriteSyncResult(c.handler.Sync())
	} else {
		c.tryStart()
	}
}

func (c *Controller) newCommand(req request) (domain.Command, bool) {
	switch req.Type {
	case PlayRequest:
		p, ok := req.Data["position"].(float64)
		if !ok {
			return nil, false
		}

		position, ok := board.ParsePosition(p)
		if !ok {
			return nil, false
		}

		command := domain.Play{
			UserID:   req.UserID,
			Position: position,
			Time:     c.timer.UTCNow(),
		}

		return command, true
	default:
		return nil, false
	}
}

func (c *Controller) observe(outgoing <-chan api.Result) {
	for result := range outgoing {
		switch data := result.(type) {
		case api.ErrorResult:
			c.send(data.UserID, data)
		default:
			c.broadcast(result)
		}
	}

	c.RLock()
	g := c.handler.Game()
	c.RUnlock()

	_ = c.store.Set(g)

	c.cleanCloseAll()
}

func (c *Controller) send(userID user.ID, result api.ErrorResult) {
	c.RLock()
	conn, ok := c.connections[userID]
	c.RUnlock()

	if ok {
		conn.WriteErrorResult(result)
	}
}

func (c *Controller) broadcast(result api.Result) {
	switch data := result.(type) {
	case api.StartedResult:
		c.forAll(func(conn *Conn) {
			conn.WriteGameStartedResult(data)
		})
	case api.MoveResult:
		c.forAll(func(conn *Conn) {
			conn.WriteMoveResult(data)
		})
	case api.EndedResult:
		c.forAll(func(conn *Conn) {
			conn.WriteGameEndedResult(data)
		})
	}
}

func (c *Controller) tryStart() {
	c.RLock()
	currentCount := len(c.connections)
	c.RUnlock()

	if currentCount < maxUserCount {
		return
	}

	if g, err := c.store.Get(c.gameID); err == nil && g.IsPlayable() {
		c.start(g)
	} else {
		c.forceCleanup()
	}
}

func (c *Controller) start(game *domain.Game) {
	c.Lock()
	defer c.Unlock()

	if c.running {
		return
	}

	c.running = true
	c.handler = api.NewHandler(game)

	go c.observe(c.handler.Loop())
}

func (c *Controller) tryCleanup() {
	c.RLock()
	running := c.running
	c.RUnlock()

	if !running {
		c.done()
	}
}

func (c *Controller) forceCleanup() {
	c.forAll(func(conn *Conn) {
		conn.CloseAsNotFound()
	})

	c.done()
}

func (c *Controller) cleanCloseAll() {
	c.forAll(func(conn *Conn) {
		conn.CloseAsDone()
	})

	c.done()
}

func (c *Controller) forAll(f func(conn *Conn)) {
	c.RLock()
	for _, conn := range c.connections {
		f(conn)
	}
	c.RUnlock()
}

func (c *Controller) create(userID user.ID, conn *Conn) {
	c.Lock()
	defer c.Unlock()
	c.connections[userID] = conn
}

func (c *Controller) delete(userID user.ID) {
	c.Lock()
	delete(c.connections, userID)
	c.Unlock()
}
