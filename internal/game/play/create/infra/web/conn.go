// Copyright 2024 Aman Agnihotri
// SPDX-License-Identifier: AGPL-3.0-only

package web

import (
	"encoding/json"
	"time"

	"ttt/internal/base/domain/game/user"
	"ttt/internal/game/play/create/api"

	ws "github.com/lesismal/nbio/nbhttp/websocket"
)

const writeTimeout = 5 * time.Second

type Conn struct {
	conn *ws.Conn
}

func NewConn(session *Session, conn *ws.Conn) *Conn {
	conn.SetSession(session)

	return &Conn{conn}
}

func (c *Conn) OnRequest(fun func(conn *Conn, ctrl *Controller, req request)) {
	c.conn.OnMessage(func(conn *ws.Conn, msgType ws.MessageType, bytes []byte) {
		if msgType != ws.TextMessage {
			c.CloseAsInvalidMessage()
		}

		var req request

		err := json.Unmarshal(bytes, &req)
		if err != nil {
			return
		}

		session, ok := conn.Session().(*Session)
		if !ok {
			return
		}

		req.UserID = session.UserID

		fun(c, session.Controller, req)
	})
}

func (c *Conn) OnClose(fun func(ctrl *Controller, userID user.ID)) {
	c.conn.OnClose(func(conn *ws.Conn, _ error) {
		session, ok := conn.Session().(*Session)
		if !ok {
			return
		}

		fun(session.Controller, session.UserID)
	})
}

func (c *Conn) WriteGameStartedResult(result api.StartedResult) {
	c.Write(Response[api.StartedResult]{Type: Start, Data: result})
}

func (c *Conn) WriteSyncResult(result api.SyncResult) {
	c.Write(Response[api.SyncResult]{Type: Sync, Data: result})
}

func (c *Conn) WriteMoveResult(result api.MoveResult) {
	c.Write(Response[api.MoveResult]{Type: Move, Data: result})
}

func (c *Conn) WriteGameEndedResult(result api.EndedResult) {
	c.Write(Response[api.EndedResult]{Type: End, Data: result})
}

func (c *Conn) WriteErrorResult(result api.ErrorResult) {
	c.Write(Response[api.ErrorResult]{Type: Error, Data: result})
}

func (c *Conn) Write(message any) {
	bytes, err := json.Marshal(message)
	if err != nil {
		return
	}

	err = c.conn.SetWriteDeadline(time.Now().Add(writeTimeout))
	if err != nil {
		return
	}

	err = c.conn.WriteMessage(ws.TextMessage, bytes)
	if err != nil {
		return
	}
}

func (c *Conn) CloseAsDone() {
	const done = 1000

	c.Close(done, "game over")
}

func (c *Conn) CloseAsInvalidMessage() {
	const forbidden = 4003

	c.Close(forbidden, "invalid message")
}

func (c *Conn) CloseAsNotFound() {
	const notFound = 4004

	c.Close(notFound, "not found")
}

func (c *Conn) Close(code int, reason string) {
	err := c.conn.SetWriteDeadline(time.Now().Add(writeTimeout))
	if err != nil {
		return
	}

	err = c.conn.WriteClose(code, reason)
	if err != nil {
		return
	}
}
