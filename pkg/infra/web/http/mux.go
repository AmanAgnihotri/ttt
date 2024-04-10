// Copyright 2024 Aman Agnihotri
// SPDX-License-Identifier: AGPL-3.0-only

package http

import (
	"context"
	"fmt"
	"net/http"
	"strings"

	"github.com/lesismal/nbio/nbhttp/websocket"

	"ttt/pkg/api/jwt"
)

type Mux struct {
	*http.ServeMux
	validator jwt.Validator
}

func NewMux(validator jwt.Validator) *Mux {
	return &Mux{http.NewServeMux(), validator}
}

func (m *Mux) Handle(pattern string, handler func(ctx *Context)) {
	m.HandleFunc(pattern, func(writer http.ResponseWriter, req *http.Request) {
		ctx := NewContext(req.Context(), req, writer)

		handler(ctx)
	})
}

func (m *Mux) HandleWithAuth(pattern string, handler func(ctx *Context)) {
	m.HandleFunc(pattern, func(writer http.ResponseWriter, req *http.Request) {
		accessToken, ok := m.getAccessToken(req)
		if !ok {
			writer.WriteHeader(http.StatusUnauthorized)

			return
		}

		claims, err := m.validator.NewClaimsFromToken(accessToken)
		if err != nil {
			writer.WriteHeader(http.StatusUnauthorized)

			return
		}

		ctxWithClaims := context.WithValue(req.Context(), claimsKey, claims)

		ctx := NewContext(ctxWithClaims, req, writer)

		handler(ctx)
	})
}

func (m *Mux) getAccessToken(req *http.Request) (string, bool) {
	const prefix = jwt.Type + " "

	header := req.Header.Get("Authorization")
	if strings.HasPrefix(header, prefix) {
		return strings.TrimPrefix(header, prefix), true
	}

	return "", false
}

var upgrader = websocket.NewUpgrader()

func (m *Mux) Upgrade(ctx *Context) (*websocket.Conn, error) {
	conn, err := upgrader.Upgrade(ctx.rw(), ctx.req(), nil)
	if err != nil {
		return nil, fmt.Errorf("error upgrading connection: %w", err)
	}

	return conn, nil
}
