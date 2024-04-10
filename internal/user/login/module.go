// Copyright 2024 Aman Agnihotri
// SPDX-License-Identifier: AGPL-3.0-only

package login

import (
	"ttt/internal/user/common"
	"ttt/internal/user/login/create"
)

func Configure(wire common.Wire) {
	create.Configure(wire)
}
