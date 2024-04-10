// Copyright 2024 Aman Agnihotri
// SPDX-License-Identifier: AGPL-3.0-only

package web

import (
	"ttt/internal/base/api/claim"
	"ttt/pkg/infra/web/http"
)

func MapRoute(router http.Router, controller Controller) {
	router.HandleWithAuth("DELETE /user", func(ctx *http.Context) {
		userID, ok := ctx.ClaimValueAsInt64(claim.UserID)
		if ok {
			controller.Handle(ctx, request{userID})
		} else {
			ctx.Unauthorized()
		}
	})
}
