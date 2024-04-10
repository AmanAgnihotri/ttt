// Copyright 2024 Aman Agnihotri
// SPDX-License-Identifier: AGPL-3.0-only

package auth

import (
	"ttt/pkg/domain/id"
)

type Token int64

func NewToken() (Token, bool) {
	if t, ok := id.NewID(); ok {
		return Token(t), true
	}

	return 0, false
}

func ParseToken(i int64) (Token, bool) {
	if id.IsValid(i) {
		return Token(i), true
	}

	return 0, false
}
