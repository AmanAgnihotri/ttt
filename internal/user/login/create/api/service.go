// Copyright 2024 Aman Agnihotri
// SPDX-License-Identifier: AGPL-3.0-only

package api

import (
	"context"

	"ttt/internal/user/common/api/auth"
	"ttt/internal/user/common/api/user"
	"ttt/internal/user/login/create/domain"
)

type Service interface {
	OkResult(ctx context.Context, user *domain.UserData) Result
	Result(event domain.Event) Result
}

func NewService(auth auth.Service, user user.Service) Service {
	return &service{auth, user}
}

type service struct {
	auth auth.Service
	user user.Service
}

func (s *service) OkResult(ctx context.Context, user *domain.UserData) Result {
	authDTO, ok := s.auth.NewDTO(user.ID, user.Auth.Token, user.UpdateTime)
	if !ok {
		return Error{}
	}

	userDTO := s.user.NewDTO(ctx, user)

	return Created{userDTO, authDTO}
}

func (s *service) Result(event domain.Event) Result {
	switch event {
	case domain.Deleted:
		return NotFound{}
	case domain.Banned:
		return Banned{}
	case domain.Locked:
		return Locked{}
	default:
		return Error{}
	}
}
