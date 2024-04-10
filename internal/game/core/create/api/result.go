// Copyright 2024 Aman Agnihotri
// SPDX-License-Identifier: AGPL-3.0-only

package api

import (
	"ttt/internal/game/common/api/auth"
	"ttt/internal/game/common/api/game"
)

type Result interface{}

type Created struct {
	Game game.DTO `json:"game"`
	Auth auth.DTO `json:"auth"`
}

type (
	InvalidAppVersion struct{}
	NotFound          struct{}
	AlreadyExists     struct{}
	Error             struct{}
)
