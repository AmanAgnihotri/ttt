// Copyright 2024 Aman Agnihotri
// SPDX-License-Identifier: AGPL-3.0-only

package domain

import (
	"time"

	"ttt/internal/base/domain/user"
	"ttt/internal/base/domain/user/auth"
)

type Command struct {
	UserID    user.ID
	AuthToken auth.Token
	Time      time.Time
}
