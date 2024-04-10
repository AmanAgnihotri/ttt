// Copyright 2024 Aman Agnihotri
// SPDX-License-Identifier: AGPL-3.0-only

package api

import (
	"context"

	"ttt/internal/base/domain/user"
	"ttt/internal/base/domain/user/auth"
	"ttt/internal/base/domain/user/guest"
	"ttt/internal/user/core/create/domain"
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
	newID, ok := h.createID(ctx)
	if !ok {
		return Error{}
	}

	newNick, ok := h.createNick(ctx)
	if !ok {
		return Error{}
	}

	newGuest, ok := guest.NewGuest()
	if !ok {
		return Error{}
	}

	newAuth, ok := auth.NewAuth(command.Time)
	if !ok {
		return Error{}
	}

	newUser := &domain.User{
		ID:         newID,
		Nick:       newNick,
		Group:      user.Normal,
		State:      user.Guest,
		Platform:   command.Platform,
		AppVersion: command.AppVersion,
		CreateTime: command.Time,
		UpdateTime: command.Time,
		ExpiryTime: nil,
		Guest:      newGuest,
		Auth:       newAuth,
		Game:       nil,
	}

	if inserted := h.store.Insert(ctx, newUser); inserted {
		return h.OkResult(ctx, newUser)
	}

	return Error{}
}

func (h *handler) createID(ctx context.Context) (user.ID, bool) {
	const maxAttempts = 3

	for range maxAttempts {
		id, ok := user.NewID()
		if ok && h.store.IsIDAvailable(ctx, id) {
			return id, true
		}
	}

	return 0, false
}

func (h *handler) createNick(ctx context.Context) (user.Nick, bool) {
	const maxAttempts = 5

	for range maxAttempts {
		newNick, ok := user.NewNick()
		if ok && h.store.IsNickAvailable(ctx, newNick) {
			return newNick, true
		}
	}

	return "", false
}
