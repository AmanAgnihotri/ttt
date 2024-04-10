// Copyright 2024 Aman Agnihotri
// SPDX-License-Identifier: AGPL-3.0-only

package app_test

import (
	"testing"

	"ttt/internal/base/domain/app"
)

type Test struct {
	Name    string
	Input   app.Version
	IsValid bool
}

func TestApp_IsValid(t *testing.T) {
	t.Parallel()

	const (
		minimumVersion = 2
		currentVersion = 5
	)

	newApp := app.NewApp(minimumVersion, currentVersion)

	tests := []Test{
		{"test same as current version", currentVersion, true},
		{"test higher than current version", currentVersion + 1, true},
		{"test same as minimum version", minimumVersion, true},
		{"test higher than minimum version", minimumVersion + 1, true},
		{"test lower than minimum version", minimumVersion - 1, false},
	}

	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			t.Parallel()

			ok := newApp.IsValid(test.Input)
			if ok != test.IsValid {
				t.Errorf(
					"app.IsValid(%v) = %v; expected %v",
					test.Input, ok, test.IsValid)
			}
		})
	}
}

func TestApp_CanUpdate(t *testing.T) {
	t.Parallel()

	const (
		minimumVersion = 2
		currentVersion = 5
	)

	newApp := app.NewApp(minimumVersion, currentVersion)

	tests := []Test{
		{"test same as current version", currentVersion, false},
		{"test higher than current version", currentVersion + 1, false},
		{"test lower than current version", currentVersion - 1, true},
	}

	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			t.Parallel()

			ok := newApp.CanUpdate(test.Input)
			if ok != test.IsValid {
				t.Errorf(
					"app.CanUpdate(%v) = %v; expected %v",
					test.Input, ok, test.IsValid)
			}
		})
	}
}
