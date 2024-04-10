// Copyright 2024 Aman Agnihotri
// SPDX-License-Identifier: AGPL-3.0-only

package api

import (
	"context"

	"ttt/internal/user/core/delete/domain"
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

	user, err := session.Get(command.UserID)
	if err != nil {
		return NotFound{}
	}

	event := user.Delete(command.Time)

	if event != domain.Deleted {
		return h.Result(event)
	}

	if err = session.Set(user); err != nil {
		return TooManyRequests{}
	}

	return h.OkResult(user)
}
