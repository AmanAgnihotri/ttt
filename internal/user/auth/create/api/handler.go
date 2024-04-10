// Copyright 2024 Aman Agnihotri
// SPDX-License-Identifier: AGPL-3.0-only

package api

import (
	"context"

	"ttt/internal/user/auth/create/domain"
)

type Handler interface {
	Handle(ctx context.Context, command domain.Command) Result
}

func NewHandler(store domain.Store, service Service) Handler {
	return &handler{store, service}
}

type handler struct {
	store domain.Store
	Service
}

func (h *handler) Handle(ctx context.Context, command domain.Command) Result {
	session := h.store.NewSession(ctx)
	defer session.End()

	user, err := session.Get(command.UserID, command.AuthToken)
	if err != nil {
		return NotFound{}
	}

	if !h.store.GetApp(ctx).IsValid(user.AppVersion) {
		return InvalidAppVersion{}
	}

	event := user.Authorise(command.Time)

	if event != domain.Authorised {
		return h.Result(event)
	}

	err = session.Set(user)
	if err != nil {
		return TooManyRequests{}
	}

	return h.OkResult(user)
}
