// Copyright 2024 Aman Agnihotri
// SPDX-License-Identifier: AGPL-3.0-only

package api

import (
	"context"

	"ttt/internal/game/user/create/domain"
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

	user, err := session.GetUser(command.UserID)
	{
		if err != nil {
			return UserNotFound{}
		}

		if !h.store.GetApp(ctx).IsValid(user.AppVersion) {
			return InvalidAppVersion{}
		}

		if !user.CanJoinGame(command.Time) {
			return AlreadyExists{}
		}
	}

	game, err := session.GetGame(command.GameCode)
	{
		if err != nil || game.IsFull() {
			return GameNotFound{}
		}

		if game.HasUser(user) {
			return AlreadyExists{}
		}
	}

	game.AddUser(user, command.Time)
	game.ExtendExpiryTime()

	user.SetGame(game)

	err = session.Set(game, user)
	if err != nil {
		return Error{}
	}

	return h.OkResult(game)
}
