// Copyright 2024 Aman Agnihotri
// SPDX-License-Identifier: AGPL-3.0-only

package user

import (
	"ttt/pkg/domain/nick"
)

type Nick string

func NewNick() (Nick, bool) {
	if n, ok := nick.NewNick(); ok {
		return Nick(n), true
	}

	return "", false
}
