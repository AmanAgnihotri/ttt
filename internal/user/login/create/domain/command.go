// Copyright 2024 Aman Agnihotri
// SPDX-License-Identifier: AGPL-3.0-only

package domain

import (
	"time"

	"ttt/internal/base/domain/app"
	"ttt/internal/base/domain/user"
	"ttt/internal/base/domain/user/guest"
)

type Command struct {
	UserID     user.ID
	Platform   app.Platform
	AppVersion app.Version
	GuestToken guest.Token
	Time       time.Time
}
