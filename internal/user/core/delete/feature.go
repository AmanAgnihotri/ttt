// Copyright 2024 Aman Agnihotri
// SPDX-License-Identifier: AGPL-3.0-only

package delete

import (
	"ttt/internal/user/common"
	"ttt/internal/user/core/delete/api"
	"ttt/internal/user/core/delete/infra/data"
	"ttt/internal/user/core/delete/infra/web"
)

func Configure(wire common.Wire) {
	controller := newController(wire)

	web.MapRoute(wire.Router(), controller)
}

func newController(wire common.Wire) web.Controller {
	store := data.NewStore(wire.DBService())
	service := api.NewService()
	handler := api.NewHandler(store, service)
	controller := web.NewController(wire.Timer(), handler)

	return controller
}
