// Copyright 2024 Aman Agnihotri
// SPDX-License-Identifier: AGPL-3.0-only

package board

const (
	Side = 3           // Side denotes NxN board
	Size = Side * Side // Size denotes NxN value
)

var (
	ranks = [][]Position{
		{0, 1, 2},
		{3, 4, 5},
		{6, 7, 8},
	}

	files = [][]Position{
		{0, 3, 6},
		{1, 4, 7},
		{2, 5, 8},
	}

	diagonals = [][]Position{
		{0, 4, 8}, // major diagonal
		{2, 4, 6}, // minor diagonal
	}
)

type Board struct {
	marks []Mark
	state State
	count byte
}

func NewBoard() *Board {
	marks := make([]Mark, Size)
	for i := range Size {
		marks[i] = Empty
	}

	return &Board{
		marks: marks,
		state: Playable,
		count: 0,
	}
}

func (b *Board) Marks() []Mark { return b.marks }
func (b *Board) State() State  { return b.state }

func (b *Board) Apply(position Position, mark Mark) bool {
	if !b.marks[position].IsEmpty() {
		return false
	}

	b.marks[position] = mark
	b.count++
	b.state = b.nextState(position)

	return true
}

func (b *Board) nextState(pos Position) State {
	const threshold = 2*Side - 1 // least number of marks before win is possible

	if b.count < threshold {
		return Playable
	}

	if b.isRankWin(pos) || b.isFileWin(pos) || b.isDiagonalWin(pos) {
		return Win
	}

	if b.count == Size {
		return Draw
	}

	return Playable
}

func (b *Board) isRankWin(position Position) bool {
	rankIndex := position / Side

	positions := ranks[rankIndex]

	return b.isStreak(positions)
}

func (b *Board) isFileWin(position Position) bool {
	fileIndex := position % Side

	positions := files[fileIndex]

	return b.isStreak(positions)
}

func (b *Board) isDiagonalWin(position Position) bool {
	rankIndex := position / Side
	fileIndex := position % Side

	if rankIndex == fileIndex {
		const major = 0
		if b.isStreak(diagonals[major]) {
			return true
		}
	}

	if rankIndex+fileIndex == Side-1 {
		const minor = 1
		if b.isStreak(diagonals[minor]) {
			return true
		}
	}

	return false
}

func (b *Board) isStreak(positions []Position) bool {
	if len(positions) == 0 {
		return false
	}

	mark := b.marks[positions[0]]
	if mark.IsEmpty() {
		return false
	}

	for _, p := range positions[1:] {
		if b.marks[p] != mark {
			return false
		}
	}

	return true
}
