// Copyright 2024 Aman Agnihotri
// SPDX-License-Identifier: AGPL-3.0-only

package api

import (
	"context"

	"ttt/internal/user/login/create/domain"
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
	if !h.store.GetApp(ctx).IsValid(command.AppVersion) {
		return InvalidAppVersion{}
	}

	return h.handle(ctx, command)
}

func (h *handler) handle(ctx context.Context, command domain.Command) Result {
	session := h.store.NewSession(ctx)
	defer session.End()

	user, err := session.Get(command.UserID, command.GuestToken)
	if err != nil {
		return NotFound{}
	}

	event := user.Login(command.Platform, command.AppVersion, command.Time)

	if event != domain.LoggedIn {
		return h.Result(event)
	}

	userData, err := session.Set(user)
	if err != nil {
		return TooManyRequests{}
	}

	return h.OkResult(ctx, userData)
}
