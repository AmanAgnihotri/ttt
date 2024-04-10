// Copyright 2024 Aman Agnihotri
// SPDX-License-Identifier: AGPL-3.0-only

package user

import (
	"ttt/internal/base"
	"ttt/internal/user/auth"
	"ttt/internal/user/common"
	"ttt/internal/user/core"
	"ttt/internal/user/login"
)

func Configure(base base.Wire) {
	wire := common.NewWire(base)

	auth.Configure(wire)
	core.Configure(wire)
	login.Configure(wire)
}
