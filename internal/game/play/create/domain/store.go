// Copyright 2024 Aman Agnihotri
// SPDX-License-Identifier: AGPL-3.0-only

package domain

import (
	"ttt/internal/base/domain/game"
)

type Store interface {
	IsActive(gameID game.ID) bool
	Get(gameID game.ID) (*Game, error)
	Set(game *Game) error
}
