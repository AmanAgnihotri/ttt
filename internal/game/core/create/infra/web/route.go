// Copyright 2024 Aman Agnihotri
// SPDX-License-Identifier: AGPL-3.0-only

package web

import (
	"ttt/internal/base/api/claim"
	"ttt/pkg/infra/web/http"
)

func MapRoute(router http.Router, controller Controller) {
	router.HandleWithAuth("POST /games", func(ctx *http.Context) {
		var req request

		if err := ctx.Body(&req); err == nil {
			userID, ok := ctx.ClaimValueAsInt64(claim.UserID)
			if ok {
				req.UserID = userID

				controller.Handle(ctx, req)
			} else {
				ctx.Unauthorized()
			}
		} else {
			ctx.BadRequest()
		}
	})
}
