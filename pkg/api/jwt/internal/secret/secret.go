// Copyright 2024 Aman Agnihotri
// SPDX-License-Identifier: AGPL-3.0-only

package secret

type Secret string

const (
	MinLength = 32
)

func Parse(s string) (Secret, bool) {
	if len(s) >= MinLength {
		return Secret(s), true
	}

	return "", false
}
