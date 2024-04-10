// Copyright 2024 Aman Agnihotri
// SPDX-License-Identifier: AGPL-3.0-only

package adjective_test

import (
	"testing"

	"ttt/pkg/domain/nick/internal/adjective"
)

func TestNewAdjective(t *testing.T) {
	t.Parallel()

	if newAdjective, ok := adjective.NewAdjective(); !ok {
		t.Errorf("adjective.NewAdjective() threw error")
	} else if !isValid(newAdjective, adjective.Values) {
		t.Errorf("invalid adjective: %s", newAdjective)
	}
}

func isValid(adjective adjective.Adjective, possibleValues []string) bool {
	adjectiveAsString := string(adjective)

	if len(adjectiveAsString) == 0 {
		return false
	}

	for _, v := range possibleValues {
		if v == adjectiveAsString {
			return true
		}
	}

	return false
}
