// Copyright 2024 Aman Agnihotri
// SPDX-License-Identifier: AGPL-3.0-only

package secret_test

import (
	"strings"
	"testing"

	"ttt/pkg/api/jwt/internal/secret"
)

func TestParse(t *testing.T) {
	type Test struct {
		Name    string
		Input   string
		Output  secret.Secret
		IsValid bool
	}

	t.Parallel()

	var (
		validSecret = strings.Repeat("x", secret.MinLength)
		longSecret  = strings.Repeat("x", secret.MinLength+1)
	)

	tests := []Test{
		{"test valid secret", validSecret, secret.Secret(validSecret), true},
		{"test long secret", longSecret, secret.Secret(longSecret), true},
		{"test short secret", "something", "", false},
		{"test empty secret", "", "", false},
	}

	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			t.Parallel()

			result, ok := secret.Parse(test.Input)
			if result != test.Output || ok != test.IsValid {
				t.Errorf(
					"secret.Parse(%v) = %v, %v; expected %v, %v",
					test.Input, result, ok, test.Output, test.IsValid)
			}
		})
	}
}
