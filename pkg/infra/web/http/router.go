// Copyright 2024 Aman Agnihotri
// SPDX-License-Identifier: AGPL-3.0-only

package http

type Router interface {
	Handle(pattern string, handler func(ctx *Context))
	HandleWithAuth(pattern string, handler func(ctx *Context))
}
