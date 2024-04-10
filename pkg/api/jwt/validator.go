// Copyright 2024 Aman Agnihotri
// SPDX-License-Identifier: AGPL-3.0-only

package jwt

type Validator interface {
	NewClaimsFromToken(token string) (Claims, error)
}
