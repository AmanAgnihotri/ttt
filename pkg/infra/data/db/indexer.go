// Copyright 2024 Aman Agnihotri
// SPDX-License-Identifier: AGPL-3.0-only

package db

import "context"

type Indexer interface {
	EnsureIndexes(ctx context.Context) error
}
