// Copyright 2024 Aman Agnihotri
// SPDX-License-Identifier: AGPL-3.0-only

package play

import (
	"ttt/internal/game/common"
	"ttt/internal/game/play/create"
)

func Configure(wire common.Wire) {
	create.Configure(wire)
}
