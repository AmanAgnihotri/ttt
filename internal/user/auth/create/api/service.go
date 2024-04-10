// Copyright 2024 Aman Agnihotri
// SPDX-License-Identifier: AGPL-3.0-only

package api

import (
	"ttt/internal/user/auth/create/domain"
	"ttt/internal/user/common/api/auth"
)

type Service interface {
	OkResult(user *domain.User) Result
	Result(event domain.Event) Result
}

func NewService(auth auth.Service) Service {
	return &service{auth}
}

type service struct {
	auth auth.Service
}

func (s *service) OkResult(user *domain.User) Result {
	authDTO, ok := s.auth.NewDTO(user.ID, user.Auth.Token, user.UpdateTime)
	if !ok {
		return Error{}
	}

	return authDTO
}

func (s *service) Result(event domain.Event) Result {
	switch event {
	case domain.TokenExpired:
		return TokenExpired{}
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
