// Copyright 2024 Aman Agnihotri
// SPDX-License-Identifier: AGPL-3.0-only

package web

import (
	"ttt/internal/base/domain/user"
	"ttt/internal/user/core/get/api"
	"ttt/internal/user/core/get/domain"
	"ttt/pkg/infra/web/http"
)

type Controller interface {
	Handle(ctx *http.Context, req request)
}

func NewController(handler api.Handler) Controller {
	return &controller{handler}
}

type controller struct {
	handler api.Handler
}

func (c *controller) Handle(ctx *http.Context, req request) {
	if query, ok := c.newQuery(req); ok {
		result := c.handler.Handle(ctx, query)

		switch data := result.(type) {
		case api.Get:
			ctx.Ok(data)
		default:
			ctx.Status(http.StatusNotFound)
		}
	} else {
		ctx.Status(http.StatusBadRequest)
	}
}

func (c *controller) newQuery(req request) (domain.Query, bool) {
	var query domain.Query

	userID, ok := user.ParseID(req.UserID)
	if !ok {
		return query, false
	}

	query = domain.Query{
		UserID: userID,
	}

	return query, true
}
