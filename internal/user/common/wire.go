// Copyright 2024 Aman Agnihotri
// SPDX-License-Identifier: AGPL-3.0-only

package common

import (
	"ttt/internal/base"
	"ttt/internal/user/common/api/auth"
	"ttt/internal/user/common/api/user"
)

type Wire interface {
	base.Wire

	AuthService() auth.Service
	UserService() user.Service
}

func NewWire(base base.Wire) Wire {
	authService := auth.NewService(base.JWTService())
	userService := user.NewService(base.AppStore())

	return &wire{base, authService, userService}
}

type wire struct {
	base.Wire

	authService auth.Service
	userService user.Service
}

func (w *wire) AuthService() auth.Service { return w.authService }
func (w *wire) UserService() user.Service { return w.userService }
