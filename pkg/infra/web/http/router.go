// Copyright 2024 Aman Agnihotri
// SPDX-License-Identifier: AGPL-3.0-only

package http

import (
	"github.com/lesismal/nbio/nbhttp/websocket"
)

type Router interface {
	Handle(pattern string, handler func(ctx *Context))
	HandleWithAuth(pattern string, handler func(ctx *Context))
	Upgrade(ctx *Context) (*websocket.Conn, error)
}
