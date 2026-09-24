// Copyright 2026 Aman Agnihotri
// SPDX-License-Identifier: AGPL-3.0-only

package board_test

import (
	"testing"

	"ttt/internal/game/common/domain/board"
)

func TestParsePosition(t *testing.T) {
	type Test struct {
		Name    string
		Input   float64
		Output  board.Position
		IsValid bool
	}

	t.Parallel()

	tests := []Test{
		{"test first position", 0, 0, true},
		{"test middle position", 4, 4, true},
		{"test last position", 8, 8, true},
		{"test negative position", -1, 0, false},
		{"test position past the board", 9, 0, false},
		{"test fractional position", 4.5, 0, false},
	}

	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			t.Parallel()

			result, ok := board.ParsePosition(test.Input)
			if result != test.Output || ok != test.IsValid {
				t.Errorf(
					"board.ParsePosition(%v) = %v, %v; expected %v, %v",
					test.Input, result, ok, test.Output, test.IsValid)
			}
		})
	}
}
