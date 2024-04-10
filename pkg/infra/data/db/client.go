// Copyright 2024 Aman Agnihotri
// SPDX-License-Identifier: AGPL-3.0-only

package db

import "context"

type Client interface {
	Ping(ctx context.Context) error
	Disconnect(ctx context.Context) error
}
