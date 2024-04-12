// Copyright 2024 Aman Agnihotri
// SPDX-License-Identifier: AGPL-3.0-only

package board

type State byte

const (
	Playable State = iota
	Win
	Draw
)

func (s State) String() string {
	switch s {
	case Playable:
		return "playable"
	case Win:
		return "win"
	case Draw:
		return "draw"
	default:
		return "unknown"
	}
}
