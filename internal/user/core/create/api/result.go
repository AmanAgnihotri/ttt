// Copyright 2024 Aman Agnihotri
// SPDX-License-Identifier: AGPL-3.0-only

package api

import (
	"ttt/internal/user/common/api/auth"
	"ttt/internal/user/common/api/user"
)

type Result interface{}

type Created struct {
	User user.DTO `json:"user"`
	Auth auth.DTO `json:"auth"`
}

type (
	InvalidAppVersion struct{}
	Error             struct{}
)
