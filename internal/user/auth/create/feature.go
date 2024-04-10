// Copyright 2024 Aman Agnihotri
// SPDX-License-Identifier: AGPL-3.0-only

package create

import (
	"ttt/internal/user/auth/create/api"
	"ttt/internal/user/auth/create/infra/data"
	"ttt/internal/user/auth/create/infra/web"
	"ttt/internal/user/common"
)

func Configure(wire common.Wire) {
	controller := newController(wire)

	web.MapRoute(wire.Router(), controller)
}

func newController(wire common.Wire) web.Controller {
	store := data.NewStore(wire.DBService(), wire.AppStore())
	service := api.NewService(wire.AuthService())
	handler := api.NewHandler(store, service)
	controller := web.NewController(wire.Timer(), handler)

	return controller
}
