// Copyright 2024 Aman Agnihotri
// SPDX-License-Identifier: AGPL-3.0-only

package domain

import (
	"time"

	"ttt/internal/base/domain/app"
)

type Command struct {
	Platform   app.Platform
	AppVersion app.Version
	Time       time.Time
}
