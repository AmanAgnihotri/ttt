// Copyright 2024 Aman Agnihotri
// SPDX-License-Identifier: AGPL-3.0-only

package app_test

import (
	"testing"

	"ttt/internal/base/domain/app"
)

func TestParsePlatform(t *testing.T) {
	type Test struct {
		Name    string
		Input   string
		Output  app.Platform
		IsValid bool
	}

	t.Parallel()

	tests := []Test{
		{"test android platform", "android", app.Android, true},
		{"test ios platform", "ios", app.IOS, true},
		{"test editor platform", "editor", app.Editor, true},
		{"test browser platform", "browser", app.Browser, true},
		{"test invalid platform", "invalid", "", false},
		{"test empty platform", "", "", false},
	}

	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			t.Parallel()

			result, ok := app.ParsePlatform(test.Input)
			if result != test.Output || ok != test.IsValid {
				t.Errorf(
					"app.ParsePlatform(%v) = %v; expected %v, %v",
					test.Input, ok, test.Output, test.IsValid)
			}
		})
	}
}
