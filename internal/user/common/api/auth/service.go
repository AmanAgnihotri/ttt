// Copyright 2024 Aman Agnihotri
// SPDX-License-Identifier: AGPL-3.0-only

package auth

import (
	"time"

	"ttt/internal/base/api/claim"
	"ttt/internal/base/domain/user"
	"ttt/internal/base/domain/user/auth"
	"ttt/internal/user/common/api/auth/refresh"
	"ttt/pkg/api/jwt"
)

const (
	lifetimeInSeconds = 3600
	lifetime          = lifetimeInSeconds * time.Second
)

type Service interface {
	NewDTO(userID user.ID, authToken auth.Token, time time.Time) (DTO, bool)
}

func NewService(jwt jwt.Service) Service {
	return &service{jwt}
}

type service struct {
	jwt.Service
}

func (s *service) NewDTO(
	userID user.ID,
	authToken auth.Token,
	time time.Time,
) (DTO, bool) {
	var dto DTO

	claims, err := s.NewClaims(time, time.Add(lifetime))
	if err != nil {
		return dto, false
	}

	claims[claim.UserID] = userID

	accessToken, err := s.NewTokenFromClaims(claims)
	if err != nil {
		return dto, false
	}

	refreshToken := refresh.NewToken(userID, authToken)

	dto = DTO{
		TokenType:        jwt.Type,
		AccessToken:      accessToken,
		ExpiresInSeconds: lifetimeInSeconds,
		RefreshToken:     refreshToken.String(),
	}

	return dto, true
}
