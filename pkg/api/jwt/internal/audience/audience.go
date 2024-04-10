// Copyright 2024 Aman Agnihotri
// SPDX-License-Identifier: AGPL-3.0-only

package audience

const (
	MinLength = 1
	MaxLength = 64
)

type Audience string

func Parse(s string) (Audience, bool) {
	if len(s) < MinLength || len(s) > MaxLength {
		return "", false
	}

	return Audience(s), true
}
