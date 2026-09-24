// Copyright 2024-2026 Aman Agnihotri
// SPDX-License-Identifier: AGPL-3.0-only

package api

import (
	"ttt/internal/user/common/api/auth"
)

type Result any

type (
	Created           = auth.DTO
	TokenExpired      struct{}
	NotFound          struct{}
	Banned            struct{}
	Locked            struct{}
	InvalidAppVersion struct{}
	TooManyRequests   struct{}
	Error             struct{}
)
