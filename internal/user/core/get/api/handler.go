// Copyright 2024 Aman Agnihotri
// SPDX-License-Identifier: AGPL-3.0-only

package api

import (
	"context"

	"ttt/internal/user/core/get/domain"
)

type Handler interface {
	Handle(ctx context.Context, query domain.Query) Result
}

func NewHandler(store domain.Store, service Service) Handler {
	return &handler{store, service}
}

type handler struct {
	store domain.Store
	Service
}

func (h *handler) Handle(ctx context.Context, query domain.Query) Result {
	if user, err := h.store.Get(ctx, query.UserID); err == nil {
		return h.NewDTO(ctx, user)
	}

	return NotFound{}
}
