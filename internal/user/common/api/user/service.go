// Copyright 2024 Aman Agnihotri
// SPDX-License-Identifier: AGPL-3.0-only

package user

import (
	"context"

	"ttt/internal/base/domain/app"
	"ttt/internal/base/domain/user"
	"ttt/internal/base/domain/user/game"
	"ttt/internal/user/common/api/user/login"
)

type Service interface {
	NewDTO(ctx context.Context, user *user.User) DTO
}

func NewService(app app.Store) Service {
	return &service{app}
}

type service struct {
	app.Store
}

func (s *service) NewDTO(ctx context.Context, user *user.User) DTO {
	return DTO{
		ID:           user.ID,
		Nick:         user.Nick,
		Group:        user.Group,
		State:        user.State,
		CreateTime:   user.CreateTime,
		UpdateTime:   user.UpdateTime,
		ExpiryTime:   user.ExpiryTime,
		LoginID:      login.NewID(user.ID, user.Guest.Token),
		CanUpdateApp: s.GetApp(ctx).CanUpdate(user.AppVersion),
		Game:         s.getGame(user),
	}
}

func (s *service) getGame(user *user.User) *game.Game {
	userGame, time := user.Game, user.UpdateTime

	if userGame != nil && userGame.ExpiryTime.After(time) {
		return userGame
	}

	return nil
}
