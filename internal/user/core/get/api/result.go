// Copyright 2024 Aman Agnihotri
// SPDX-License-Identifier: AGPL-3.0-only

package api

import (
	"ttt/internal/user/common/api/user"
)

type Result interface{}

type (
	Get      = user.DTO
	NotFound struct{}
)
