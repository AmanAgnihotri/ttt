// Copyright 2024 Aman Agnihotri
// SPDX-License-Identifier: AGPL-3.0-only

package api

import (
	"context"

	"ttt/internal/base/domain/game"
	"ttt/internal/game/core/create/domain"
	"ttt/internal/game/core/create/domain/entity"
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
	{
		if err != nil {
			return NotFound{}
		}

		if !h.store.GetApp(ctx).IsValid(user.AppVersion) {
			return InvalidAppVersion{}
		}

		if !user.CanCreateGame(command.Time) {
			return AlreadyExists{}
		}
	}

	newID, ok := h.createID(ctx)
	if !ok {
		return Error{}
	}

	newCode, ok := h.createCode(ctx)
	if !ok {
		return Error{}
	}

	newGame := entity.NewGame(newID, newCode, command.Time, user, command.Side)

	user.SetGame(newGame)

	err = session.Set(newGame, user)
	if err != nil {
		return Error{}
	}

	return h.OkResult(newGame)
}

func (h *handler) createID(ctx context.Context) (game.ID, bool) {
	const maxAttempts = 3

	for range maxAttempts {
		id, ok := game.NewID()
		if ok && h.store.IsIDAvailable(ctx, id) {
			return id, true
		}
	}

	return 0, false
}

func (h *handler) createCode(ctx context.Context) (game.Code, bool) {
	const maxAttempts = 5

	for range maxAttempts {
		c, ok := game.NewCode()
		if ok && h.store.IsCodeAvailable(ctx, c) {
			return c, true
		}
	}

	return "", false
}
