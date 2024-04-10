// Copyright 2024 Aman Agnihotri
// SPDX-License-Identifier: AGPL-3.0-only

package get

import (
	"ttt/internal/user/common"
	"ttt/internal/user/core/get/api"
	"ttt/internal/user/core/get/infra/data"
	"ttt/internal/user/core/get/infra/web"
)

func Configure(wire common.Wire) {
	controller := newController(wire)

	web.MapRoute(wire.Router(), controller)
}

func newController(wire common.Wire) web.Controller {
	store := data.NewStore(wire.DBService())
	handler := api.NewHandler(store, wire.UserService())
	controller := web.NewController(handler)

	return controller
}
