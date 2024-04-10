// Copyright 2024 Aman Agnihotri
// SPDX-License-Identifier: AGPL-3.0-only

package api

import (
	"ttt/internal/game/common/api/auth"
	"ttt/internal/game/common/api/game"
	"ttt/internal/game/core/create/domain/entity"
)

type Service interface {
	OkResult(game *entity.Game) Result
}

func NewService(auth auth.Service) Service {
	return &service{auth}
}

type service struct {
	auth auth.Service
}

func (s *service) OkResult(game *entity.Game) Result {
	userID := game.Users[0].ID

	authDTO, ok := s.auth.NewDTO(game.ID, userID, game.CreateTime)
	if !ok {
		return Error{}
	}

	gameDTO := newDTO(game)

	return Created{gameDTO, authDTO}
}

func newDTO(g *entity.Game) game.DTO {
	return game.DTO{
		Code:       g.Code,
		CreateTime: g.CreateTime,
		ExpiryTime: g.ExpiryTime,
	}
}
