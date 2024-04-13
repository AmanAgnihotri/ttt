// Copyright 2024 Aman Agnihotri
// SPDX-License-Identifier: AGPL-3.0-only

package web

import (
	"sync"

	"ttt/internal/base/domain/game"
	"ttt/internal/base/domain/game/user"
	"ttt/internal/game/play/create/domain"
	"ttt/pkg/infra/web/time"

	ws "github.com/lesismal/nbio/nbhttp/websocket"
)

type Hub interface {
	Handle(gameID game.ID, userID user.ID, conn *ws.Conn)
}

func NewHub(timer time.Timer, store domain.Store) Hub {
	return &hub{
		timer:       timer,
		store:       store,
		controllers: make(map[game.ID]*Controller),
		RWMutex:     sync.RWMutex{},
	}
}

type hub struct {
	timer       time.Timer
	store       domain.Store
	controllers map[game.ID]*Controller
	sync.RWMutex
}

func (h *hub) Handle(gameID game.ID, userID user.ID, conn *ws.Conn) {
	h.RLock()
	ctrl, exists := h.controllers[gameID]
	h.RUnlock()

	if !exists {
		if !h.store.IsActive(gameID) {
			const notFound = 4004

			err := conn.WriteClose(notFound, "not found")
			if err != nil {
				return
			}
		}

		ctrl = h.create(gameID)
	}

	session := &Session{userID, ctrl}
	ctrl.Handle(userID, NewConn(session, conn))
}

func (h *hub) create(gameID game.ID) *Controller {
	h.Lock()
	defer h.Unlock()

	ctrl, exists := h.controllers[gameID]
	if !exists {
		done := func() { h.delete(gameID) }
		ctrl = NewController(h.timer, h.store, gameID, done)
		h.controllers[gameID] = ctrl
	}

	return ctrl
}

func (h *hub) delete(gameID game.ID) {
	h.Lock()
	defer h.Unlock()
	delete(h.controllers, gameID)
}
