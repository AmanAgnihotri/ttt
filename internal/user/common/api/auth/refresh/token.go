// Copyright 2024 Aman Agnihotri
// SPDX-License-Identifier: AGPL-3.0-only

package refresh

import (
	"fmt"
	"strconv"

	"ttt/internal/base/domain/user"
	"ttt/internal/base/domain/user/auth"
)

type Token string

func NewToken(userID user.ID, authToken auth.Token) Token {
	return Token(fmt.Sprintf("%d%d", authToken, userID))
}

func Parse(value string) (user.ID, auth.Token, bool) {
	const tokenLength = 32
	if len(value) != tokenLength {
		return 0, 0, false
	}

	part1 := value[:16]
	part2 := value[16:]

	token, err := strconv.ParseInt(part1, 10, 64)
	if err != nil {
		return 0, 0, false
	}

	userID, err := strconv.ParseInt(part2, 10, 64)
	if err != nil {
		return 0, 0, false
	}

	return user.ID(userID), auth.Token(token), true
}

func (t Token) Split() (user.ID, auth.Token, bool) {
	return Parse(t.String())
}

func (t Token) String() string {
	return string(t)
}
