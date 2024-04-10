// Copyright 2024 Aman Agnihotri
// SPDX-License-Identifier: AGPL-3.0-only

package domain

import "context"

type Store interface {
	NewSession(ctx context.Context) Session
}
