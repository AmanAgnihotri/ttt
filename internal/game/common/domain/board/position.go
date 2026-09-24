// Copyright 2024-2026 Aman Agnihotri
// SPDX-License-Identifier: AGPL-3.0-only

package board

import (
	"math"
)

type Position byte

func ParsePosition(s float64) (Position, bool) {
	if s < 0 || s >= Size || s != math.Trunc(s) {
		return 0, false
	}

	return Position(s), true
}
