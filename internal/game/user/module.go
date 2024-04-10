// Copyright 2024 Aman Agnihotri
// SPDX-License-Identifier: AGPL-3.0-only

package user

import (
	"ttt/internal/game/common"
	"ttt/internal/game/user/create"
)

func Configure(wire common.Wire) {
	create.Configure(wire)
}
