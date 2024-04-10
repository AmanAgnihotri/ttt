// Copyright 2024 Aman Agnihotri
// SPDX-License-Identifier: AGPL-3.0-only

package domain

import (
	"ttt/internal/base/domain/game/user"
	"ttt/internal/game/core/create/domain/entity"
)

type Session interface {
	Get(userID user.ID) (*entity.User, error)
	Set(game *entity.Game, user *entity.User) error
	End()
}
