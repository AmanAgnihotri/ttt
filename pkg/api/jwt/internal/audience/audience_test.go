// Copyright 2024 Aman Agnihotri
// SPDX-License-Identifier: AGPL-3.0-only

package audience_test

import (
	"strings"
	"testing"

	"ttt/pkg/api/jwt/internal/audience"
)

func TestParse(t *testing.T) {
	type Test struct {
		Name    string
		Input   string
		Output  audience.Audience
		IsValid bool
	}

	t.Parallel()

	uri := "https://example.com"
	invalid := strings.Repeat("x", audience.MaxLength+1)

	tests := []Test{
		{"test valid audience", "t", audience.Audience("t"), true},
		{"test another valid audience", uri, audience.Audience(uri), true},
		{"test very long audience", invalid, "", false},
		{"test empty audience", "", "", false},
	}

	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			t.Parallel()

			result, ok := audience.Parse(test.Input)
			if result != test.Output || ok != test.IsValid {
				t.Errorf(
					"audience.Parse(%v) = %v, %v; expected %v, %v",
					test.Input, result, ok, test.Output, test.IsValid)
			}
		})
	}
}
