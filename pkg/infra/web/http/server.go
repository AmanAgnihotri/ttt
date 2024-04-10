// Copyright 2024 Aman Agnihotri
// SPDX-License-Identifier: AGPL-3.0-only

package http

import (
	"context"
	"fmt"

	"github.com/lesismal/nbio/nbhttp"
)

type Server struct {
	engine *nbhttp.Engine
}

func NewServer(name, address string, maxLoad int, mux *Mux) *Server {
	return &Server{nbhttp.NewEngine(nbhttp.Config{
		Name:                    name,
		Network:                 "tcp",
		Addrs:                   []string{address},
		MaxLoad:                 maxLoad,
		ReleaseWebsocketPayload: true,
		Handler:                 mux,
		IOMod:                   nbhttp.IOModMixed,
	})}
}

func (s *Server) Start() error {
	if err := s.engine.Start(); err != nil {
		return fmt.Errorf("error starting server: %w", err)
	}

	return nil
}

func (s *Server) Stop(ctx context.Context) error {
	if err := s.engine.Shutdown(ctx); err != nil {
		return fmt.Errorf("error stopping server: %w", err)
	}

	return nil
}
