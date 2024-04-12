// Copyright 2024 Aman Agnihotri
// SPDX-License-Identifier: AGPL-3.0-only

package login

import (
	"fmt"
	"strconv"

	"ttt/internal/base/domain/user"
	"ttt/internal/base/domain/user/guest"
)

type ID string

func NewID(userID user.ID, guestToken guest.Token) ID {
	return ID(fmt.Sprintf("%d%d", guestToken, userID))
}

func Parse(value string) (user.ID, guest.Token, bool) {
	const idLength = 32
	if len(value) != idLength {
		return 0, 0, false
	}

	part1 := value[:16]
	part2 := value[16:]

	guestID, err := strconv.ParseInt(part2, 10, 64)
	if err != nil {
		return 0, 0, false
	}

	guestToken, err := strconv.ParseInt(part1, 10, 64)
	if err != nil {
		return 0, 0, false
	}

	return user.ID(guestID), guest.Token(guestToken), true
}

func (i ID) Split() (user.ID, guest.Token, bool) {
	return Parse(i.String())
}

func (i ID) String() string {
	return string(i)
}
