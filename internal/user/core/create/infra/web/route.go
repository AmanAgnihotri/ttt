// Copyright 2024 Aman Agnihotri
// SPDX-License-Identifier: AGPL-3.0-only

package web

import (
	"ttt/pkg/infra/web/http"
)

func MapRoute(router http.Router, controller Controller) {
	router.Handle("POST /users", func(ctx *http.Context) {
		var req request

		if err := ctx.Body(&req); err == nil {
			controller.Handle(ctx, req)
		} else {
			ctx.BadRequest()
		}
	})
}
