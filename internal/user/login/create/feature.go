// Copyright 2024 Aman Agnihotri
// SPDX-License-Identifier: AGPL-3.0-only

package create

import (
	"ttt/internal/user/common"
	"ttt/internal/user/login/create/api"
	"ttt/internal/user/login/create/infra/data"
	"ttt/internal/user/login/create/infra/web"
)

func Configure(wire common.Wire) {
	controller := newController(wire)

	web.MapRoute(wire.Router(), controller)
}

func newController(wire common.Wire) web.Controller {
	store := data.NewStore(wire.DBService(), wire.AppStore())
	service := api.NewService(wire.AuthService(), wire.UserService())
	handler := api.NewHandler(store, service)
	controller := web.NewController(wire.Timer(), handler)

	return controller
}
