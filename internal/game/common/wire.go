// Copyright 2024 Aman Agnihotri
// SPDX-License-Identifier: AGPL-3.0-only

package common

import (
	"ttt/internal/base"
	"ttt/internal/game/common/api/auth"
)

type Wire interface {
	base.Wire

	AuthService() auth.Service
}

func NewWire(base base.Wire) Wire {
	authService := auth.NewService(base.JWTService())

	return &wire{base, authService}
}

type wire struct {
	base.Wire

	authService auth.Service
}

func (w *wire) AuthService() auth.Service { return w.authService }
