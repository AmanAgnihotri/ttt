// Copyright 2024 Aman Agnihotri
// SPDX-License-Identifier: AGPL-3.0-only

package core

import (
	"ttt/internal/user/common"
	"ttt/internal/user/core/create"
	"ttt/internal/user/core/delete"
	"ttt/internal/user/core/get"
)

func Configure(wire common.Wire) {
	create.Configure(wire)
	delete.Configure(wire)
	get.Configure(wire)
}
