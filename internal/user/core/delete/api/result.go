// Copyright 2024 Aman Agnihotri
// SPDX-License-Identifier: AGPL-3.0-only

package api

import "time"

type Result interface{}

type Deleted struct {
	ExpiryTime time.Time
}

type (
	NotFound        struct{}
	Banned          struct{}
	Locked          struct{}
	TooManyRequests struct{}
	Error           struct{}
)
