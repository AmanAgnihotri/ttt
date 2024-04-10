// Copyright 2024 Aman Agnihotri
// SPDX-License-Identifier: AGPL-3.0-only

package game

import (
	"ttt/pkg/domain/id"
)

type ID int64

func NewID() (ID, bool) {
	if i, ok := id.NewID(); ok {
		return ID(i), true
	}

	return 0, false
}

func ParseID(i int64) (ID, bool) {
	if id.IsValid(i) {
		return ID(i), true
	}

	return 0, false
}
