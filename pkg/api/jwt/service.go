// Copyright 2024 Aman Agnihotri
// SPDX-License-Identifier: AGPL-3.0-only

package jwt

import "time"

type Service interface {
	NewClaims(time, expiryTime time.Time) (Claims, error)
	NewTokenFromClaims(claims Claims) (string, error)
}
