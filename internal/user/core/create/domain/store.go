// Copyright 2024 Aman Agnihotri
// SPDX-License-Identifier: AGPL-3.0-only

package domain

import (
	"context"

	"ttt/internal/base/domain/app"
	"ttt/internal/base/domain/user"
)

type Store interface {
	app.Store
	IsIDAvailable(ctx context.Context, userID user.ID) bool
	IsNickAvailable(ctx context.Context, nick user.Nick) bool
	Insert(ctx context.Context, user *User) bool
}
