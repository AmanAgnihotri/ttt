// Copyright 2024 Aman Agnihotri
// SPDX-License-Identifier: AGPL-3.0-only

package issuer

const (
	MinLength = 1
	MaxLength = 64
)

type Issuer string

func Parse(s string) (Issuer, bool) {
	if len(s) < MinLength || len(s) > MaxLength {
		return "", false
	}

	return Issuer(s), true
}
