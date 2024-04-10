// Copyright 2024 Aman Agnihotri
// SPDX-License-Identifier: AGPL-3.0-only

package app_test

import (
	"testing"

	"ttt/internal/base/domain/app"
)

func TestParseVersion(t *testing.T) {
	type Test struct {
		Name    string
		Input   int32
		Output  app.Version
		IsValid bool
	}

	t.Parallel()

	tests := []Test{
		{"test positive version", 1, app.Version(1), true},
		{"test negative version", -1, 0, false},
		{"test zero version", 0, 0, false},
	}

	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			t.Parallel()

			result, ok := app.ParseVersion(test.Input)
			if result != test.Output || ok != test.IsValid {
				t.Errorf(
					"app.ParseVersion(%v) = %v, %v; expected %v, %v",
					test.Input, result, ok, test.Output, test.IsValid)
			}
		})
	}
}
