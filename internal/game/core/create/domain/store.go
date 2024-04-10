// Copyright 2024 Aman Agnihotri
// SPDX-License-Identifier: AGPL-3.0-only

package domain

import (
	"context"

	"ttt/internal/base/domain/app"
	"ttt/internal/base/domain/game"
)

type Store interface {
	app.Store
	IsIDAvailable(ctx context.Context, gameID game.ID) bool
	IsCodeAvailable(ctx context.Context, gameCode game.Code) bool
	NewSession(ctx context.Context) Session
}
