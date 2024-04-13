// Copyright 2024 Aman Agnihotri
// SPDX-License-Identifier: AGPL-3.0-only

package domain

type Event byte

const (
	Unplayable Event = iota
	OutOfTurnPlay
	AlreadyOccupied
	Applied
	Completed
)
