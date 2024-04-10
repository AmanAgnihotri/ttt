// Copyright 2024 Aman Agnihotri
// SPDX-License-Identifier: AGPL-3.0-only

package issuer_test

import (
	"strings"
	"testing"

	"ttt/pkg/api/jwt/internal/issuer"
)

func TestParse(t *testing.T) {
	type Test struct {
		Name    string
		Input   string
		Output  issuer.Issuer
		IsValid bool
	}

	t.Parallel()

	uri := "https://example.com"
	invalid := strings.Repeat("x", issuer.MaxLength+1)

	tests := []Test{
		{"test valid issuer", "u", issuer.Issuer("u"), true},
		{"test another valid issuer", uri, issuer.Issuer(uri), true},
		{"test very long issuer", invalid, "", false},
		{"test empty issuer", "", "", false},
	}

	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			t.Parallel()

			result, ok := issuer.Parse(test.Input)
			if result != test.Output || ok != test.IsValid {
				t.Errorf(
					"issuer.Parse(%v) = %v, %v; expected %v, %v",
					test.Input, result, ok, test.Output, test.IsValid)
			}
		})
	}
}
