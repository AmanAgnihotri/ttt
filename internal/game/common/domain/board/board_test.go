// Copyright 2026 Aman Agnihotri
// SPDX-License-Identifier: AGPL-3.0-only

package board_test

import (
	"slices"
	"testing"

	"ttt/internal/game/common/domain/board"
)

func TestNewBoard(t *testing.T) {
	t.Parallel()

	newBoard := board.NewBoard()

	if state := newBoard.State(); state != board.Playable {
		t.Errorf("state = %v; expected %v", state, board.Playable)
	}

	for _, mark := range newBoard.Marks() {
		if !mark.IsEmpty() {
			t.Errorf("mark = %c; expected %c", mark, board.Empty)
		}
	}
}

func TestWins(t *testing.T) {
	t.Parallel()

	streaks := [][]board.Position{
		{0, 1, 2}, {3, 4, 5}, {6, 7, 8},
		{0, 3, 6}, {1, 4, 7}, {2, 5, 8},
		{0, 4, 8}, {2, 4, 6},
	}

	for _, streak := range streaks {
		newBoard := board.NewBoard()
		empty := emptyPositions(streak)

		newBoard.Apply(streak[0], board.XMark)
		newBoard.Apply(empty[0], board.OMark)
		newBoard.Apply(streak[1], board.XMark)
		newBoard.Apply(empty[1], board.OMark)

		if state := newBoard.State(); state != board.Playable {
			t.Errorf("state before %v = %v; expected %v",
				streak, state, board.Playable)
		}

		newBoard.Apply(streak[2], board.XMark)

		if state := newBoard.State(); state != board.Win {
			t.Errorf("state after %v = %v; expected %v",
				streak, state, board.Win)
		}
	}
}

func TestDraw(t *testing.T) {
	t.Parallel()

	newBoard := board.NewBoard()
	marks := []board.Mark{board.XMark, board.OMark}

	for i, position := range []board.Position{0, 1, 2, 4, 3, 5, 7, 6, 8} {
		newBoard.Apply(position, marks[i%2])
	}

	if state := newBoard.State(); state != board.Draw {
		t.Errorf("state = %v; expected %v", state, board.Draw)
	}
}

func TestOccupiedPosition(t *testing.T) {
	t.Parallel()

	newBoard := board.NewBoard()

	if ok := newBoard.Apply(4, board.XMark); !ok {
		t.Fatalf("Apply(4, x) = %v; expected true", ok)
	}

	if ok := newBoard.Apply(4, board.OMark); ok {
		t.Errorf("Apply(4, o) = %v; expected false", ok)
	}
}

func emptyPositions(streak []board.Position) []board.Position {
	empty := make([]board.Position, 0, board.Size)

	for position := range board.Position(board.Size) {
		if !slices.Contains(streak, position) {
			empty = append(empty, position)
		}
	}

	return empty
}
