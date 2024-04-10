// Copyright 2024 Aman Agnihotri
// SPDX-License-Identifier: AGPL-3.0-only

package api

import (
	"ttt/internal/user/core/delete/domain"
)

type Service interface {
	OkResult(user *domain.User) Result
	Result(event domain.Event) Result
}

func NewService() Service {
	return &service{}
}

type service struct{}

func (s *service) OkResult(user *domain.User) Result {
	return Deleted{user.ExpiryTime}
}

func (s *service) Result(event domain.Event) Result {
	switch event {
	case domain.AlreadyDeleted:
		return NotFound{}
	case domain.Banned:
		return Banned{}
	case domain.Locked:
		return Locked{}
	default:
		return Error{}
	}
}
