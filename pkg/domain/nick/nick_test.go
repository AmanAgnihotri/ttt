// Copyright 2024 Aman Agnihotri
// SPDX-License-Identifier: AGPL-3.0-only

package nick_test

import (
	"testing"
	"unicode"

	"ttt/pkg/domain/nick"
)

func TestNewNick(t *testing.T) {
	t.Parallel()

	if newNick, ok := nick.NewNick(); !ok {
		t.Errorf("nick.NewNick() threw error")
	} else if !isValid(newNick) {
		t.Errorf("invalid nick: %s", newNick)
	}
}

func isValid(nick string) bool {
	const minLength = 3

	if len(nick) <= minLength {
		return false
	}

	var nounStartIndex, tagStartIndex int

	if nounStartIndex = getNounStartIndex(nick); nounStartIndex <= 0 {
		return false
	}

	if tagStartIndex = getTagStartIndex(nick); tagStartIndex <= 0 {
		return false
	}

	if adjective := nick[:nounStartIndex]; !isValidWord(adjective) {
		return false
	}

	if noun := nick[nounStartIndex:tagStartIndex]; !isValidWord(noun) {
		return false
	}

	if tag := nick[tagStartIndex:]; !isValidTag(tag) {
		return false
	}

	return true
}

func getNounStartIndex(s string) int {
	for i, r := range s {
		if unicode.IsUpper(r) && i != 0 {
			return i
		}
	}

	return -1
}

func getTagStartIndex(s string) int {
	for i, r := range s {
		if unicode.IsDigit(r) {
			return i
		}
	}

	return -1
}

func isValidWord(s string) bool {
	if !unicode.IsUpper(rune(s[0])) {
		return false
	}

	for _, r := range s {
		if !unicode.IsLetter(r) {
			return false
		}
	}

	return true
}

func isValidTag(s string) bool {
	for _, r := range s {
		if !unicode.IsDigit(r) {
			return false
		}
	}

	return true
}
