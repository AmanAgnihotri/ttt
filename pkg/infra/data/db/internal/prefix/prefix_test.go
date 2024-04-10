// Copyright 2024 Aman Agnihotri
// SPDX-License-Identifier: AGPL-3.0-only

package prefix_test

import (
	"strings"
	"testing"

	"ttt/pkg/infra/data/db/internal/prefix"
)

func TestParse(t *testing.T) {
	type Test struct {
		Name    string
		Input   string
		Output  prefix.Prefix
		IsValid bool
	}

	t.Parallel()

	longPrefix := strings.Repeat("x", prefix.MaxLength+1)

	tests := []Test{
		{"test dev prefix", "dev", prefix.Prefix("dev"), true},
		{"test staging prefix", "staging", prefix.Prefix("staging"), true},
		{"test production prefix", "prod", prefix.Prefix("prod"), true},
		{"test single character prefix", "x", prefix.Prefix("x"), true},
		{"test long prefix", longPrefix, "", false},
		{"test invalid prefix (space)", " ", "", false},
		{"test invalid prefix (special chars)", "$%#", "", false},
		{"test empty prefix", "", "", false},
	}

	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			t.Parallel()

			result, ok := prefix.Parse(test.Input)
			if result != test.Output || ok != test.IsValid {
				t.Errorf(
					"prefix.Parse(%v) = %v, %v; expected %v, %v",
					test.Input, result, ok, test.Output, test.IsValid)
			}
		})
	}
}
