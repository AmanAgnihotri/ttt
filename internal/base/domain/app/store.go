// Copyright 2024 Aman Agnihotri
// SPDX-License-Identifier: AGPL-3.0-only

package app

import "context"

type Store interface {
	GetApp(ctx context.Context) *App
}
