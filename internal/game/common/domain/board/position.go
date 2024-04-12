// Copyright 2024 Aman Agnihotri
// SPDX-License-Identifier: AGPL-3.0-only

package board

type Position byte

func ParsePosition(s float64) (Position, bool) {
	if s < 0 || s > Size {
		return 0, false
	}

	return Position(s), true
}
