// Copyright 2024 Aman Agnihotri
// SPDX-License-Identifier: AGPL-3.0-only

package create

import (
	"ttt/internal/game/common"
	"ttt/internal/game/play/create/infra/data"
	"ttt/internal/game/play/create/infra/web"
)

func Configure(wire common.Wire) {
	hub := newHub(wire)

	web.MapRoute(wire.Router(), hub)
}

func newHub(wire common.Wire) web.Hub {
	store := data.NewStore(wire.DBService())
	hub := web.NewHub(wire.Timer(), store)

	return hub
}
