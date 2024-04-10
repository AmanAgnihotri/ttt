// Copyright 2024 Aman Agnihotri
// SPDX-License-Identifier: AGPL-3.0-only

package domain

import (
	"context"

	"ttt/internal/base/domain/app"
)

type Store interface {
	app.Store
	NewSession(ctx context.Context) Session
}
