// Copyright 2024 Aman Agnihotri
// SPDX-License-Identifier: AGPL-3.0-only

package auth

import (
	"time"

	"ttt/internal/base/api/claim"
	"ttt/internal/base/domain/game"
	"ttt/internal/base/domain/game/user"
	"ttt/pkg/api/jwt"
)

const (
	lifetimeInSeconds = 2400
	lifetime          = lifetimeInSeconds * time.Second
)

type Service interface {
	NewDTO(gameID game.ID, userID user.ID, now time.Time) (DTO, bool)
}

func NewService(jwtService jwt.Service) Service {
	return &service{jwtService}
}

type service struct {
	jwt.Service
}

func (s *service) NewDTO(
	gameID game.ID,
	userID user.ID,
	now time.Time,
) (DTO, bool) {
	var dto DTO

	claims, err := s.NewClaims(now, now.Add(lifetime))
	if err != nil {
		return dto, false
	}

	claims[claim.GameID] = gameID
	claims[claim.PlayerID] = userID

	accessToken, err := s.NewTokenFromClaims(claims)
	if err != nil {
		return dto, false
	}

	dto = DTO{
		TokenType:        jwt.Type,
		AccessToken:      accessToken,
		ExpiresInSeconds: lifetimeInSeconds,
	}

	return dto, true
}
