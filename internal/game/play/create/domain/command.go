// Copyright 2024-2026 Aman Agnihotri
// SPDX-License-Identifier: AGPL-3.0-only

package domain

import (
	"time"

	"ttt/internal/base/domain/game/user"
	"ttt/internal/game/common/domain/board"
)

type Command any

type Play struct {
	UserID   user.ID
	Position board.Position
	Time     time.Time
}
