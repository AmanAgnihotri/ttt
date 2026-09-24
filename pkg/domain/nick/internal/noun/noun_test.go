// Copyright 2024-2026 Aman Agnihotri
// SPDX-License-Identifier: AGPL-3.0-only

package noun_test

import (
	"slices"
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

	return slices.Contains(possibleValues, nounAsString)
}
