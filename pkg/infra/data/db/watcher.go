// Copyright 2024 Aman Agnihotri
// SPDX-License-Identifier: AGPL-3.0-only

package db

import "context"

type Watcher interface {
	Watch(ctx context.Context) error
}
