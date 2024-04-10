// Copyright 2024 Aman Agnihotri
// SPDX-License-Identifier: AGPL-3.0-only

package domain

import (
	"ttt/internal/base/domain/game"
	"ttt/internal/base/domain/game/user"
	"ttt/internal/game/user/create/domain/entity"
)

type Session interface {
	GetUser(userID user.ID) (*entity.User, error)
	GetGame(gameCode game.Code) (*entity.Game, error)
	Set(game *entity.Game, user *entity.User) error
	End()
}
