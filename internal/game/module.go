// Copyright 2024 Aman Agnihotri
// SPDX-License-Identifier: AGPL-3.0-only

package game

import (
	"ttt/internal/base"
	"ttt/internal/game/common"
	"ttt/internal/game/core"
	"ttt/internal/game/play"
	"ttt/internal/game/user"
)

func Configure(base base.Wire) {
	wire := common.NewWire(base)

	core.Configure(wire)
	play.Configure(wire)
	user.Configure(wire)
}
