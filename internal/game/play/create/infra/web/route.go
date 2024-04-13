// Copyright 2024 Aman Agnihotri
// SPDX-License-Identifier: AGPL-3.0-only

package web

import (
	"ttt/internal/base/api/claim"
	"ttt/internal/base/domain/game"
	"ttt/internal/base/domain/game/user"
	"ttt/pkg/infra/web/http"
)

func MapRoute(router http.Router, hub Hub) {
	router.HandleWithAuth("GET /gameplay", func(ctx *http.Context) {
		if gameID, userID, ok := gameAndUserIDs(ctx); ok {
			if conn, err := router.Upgrade(ctx); err == nil {
				hub.Handle(gameID, userID, conn)
			} else {
				ctx.Status(http.StatusServiceUnavailable)
			}
		} else {
			ctx.Unauthorized()
		}
	})
}

func gameAndUserIDs(ctx *http.Context) (game.ID, user.ID, bool) {
	gameID, ok := ctx.ClaimValueAsInt64(claim.GameID)
	if !ok {
		return 0, 0, false
	}

	userID, ok := ctx.ClaimValueAsInt64(claim.PlayerID)
	if !ok {
		return 0, 0, false
	}

	return game.ID(gameID), user.ID(userID), true
}
