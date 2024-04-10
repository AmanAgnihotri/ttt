// Copyright 2024 Aman Agnihotri
// SPDX-License-Identifier: AGPL-3.0-only

package auth

import (
	"ttt/internal/user/auth/create"
	"ttt/internal/user/common"
)

func Configure(wire common.Wire) {
	create.Configure(wire)
}
