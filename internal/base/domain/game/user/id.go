// Copyright 2024 Aman Agnihotri
// SPDX-License-Identifier: AGPL-3.0-only

package user

import (
	"ttt/pkg/domain/id"
)

type ID int64

func ParseID(i int64) (ID, bool) {
	if id.IsValid(i) {
		return ID(i), true
	}

	return 0, false
}
