// Copyright 2024 Aman Agnihotri
// SPDX-License-Identifier: AGPL-3.0-only

package domain

import (
	"context"

	"ttt/internal/base/domain/user"
)

type Store interface {
	Get(ctx context.Context, userID user.ID) (*user.User, error)
}
