// Copyright 2024 Aman Agnihotri
// SPDX-License-Identifier: AGPL-3.0-only

package web

import (
	"ttt/internal/base/domain/game"
	"ttt/internal/base/domain/game/user"
	"ttt/internal/game/user/create/api"
	"ttt/internal/game/user/create/domain"
	"ttt/pkg/infra/web/http"
	"ttt/pkg/infra/web/time"
)

type Controller interface {
	Handle(ctx *http.Context, req request)
}

func NewController(timer time.Timer, handler api.Handler) Controller {
	mapping := map[api.Result]http.StatusCode{
		api.InvalidAppVersion{}: http.StatusUpgradeRequired,
		api.UserNotFound{}:      http.StatusNotFound,
		api.GameNotFound{}:      http.StatusGone,
		api.AlreadyExists{}:     http.StatusForbidden,
		api.Error{}:             http.StatusInternalServerError,
	}

	return &controller{timer, handler, mapping}
}

type controller struct {
	timer   time.Timer
	handler api.Handler
	mapping map[api.Result]http.StatusCode
}

func (c *controller) Handle(ctx *http.Context, req request) {
	if command, ok := c.newCommand(req); ok {
		result := c.handler.Handle(ctx, command)

		switch data := result.(type) {
		case api.Created:
			ctx.Ok(data)
		default:
			if statusCode, exists := c.mapping[result]; exists {
				ctx.Status(statusCode)
			} else {
				ctx.Status(http.StatusNotImplemented)
			}
		}
	} else {
		ctx.Status(http.StatusBadRequest)
	}
}

func (c *controller) newCommand(req request) (domain.Command, bool) {
	var command domain.Command

	userID, ok := user.ParseID(req.UserID)
	if !ok {
		return command, false
	}

	gameCode, ok := game.ParseCode(req.GameCode)
	if !ok {
		return command, false
	}

	command = domain.Command{
		UserID:   userID,
		GameCode: gameCode,
		Time:     c.timer.UTCNow(),
	}

	return command, true
}
