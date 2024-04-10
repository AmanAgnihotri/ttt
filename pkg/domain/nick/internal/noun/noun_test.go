// Copyright 2024 Aman Agnihotri
// SPDX-License-Identifier: AGPL-3.0-only

package noun_test

import (
	"testing"

	"ttt/pkg/domain/nick/internal/noun"
)

func TestNewNoun(t *testing.T) {
	t.Parallel()

	if newNoun, ok := noun.NewNoun(); !ok {
		t.Errorf("noun.NewNoun() threw error")
	} else if !isValid(newNoun, noun.Values) {
		t.Errorf("invalid noun: %s", newNoun)
	}
}

func isValid(noun noun.Noun, possibleValues []string) bool {
	nounAsString := string(noun)

	if len(nounAsString) == 0 {
		return false
	}

	for _, v := range possibleValues {
		if v == nounAsString {
			return true
		}
	}

	return false
}
