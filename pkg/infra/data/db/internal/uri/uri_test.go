// Copyright 2024 Aman Agnihotri
// SPDX-License-Identifier: AGPL-3.0-only

package uri_test

import (
	"strings"
	"testing"

	"ttt/pkg/infra/data/db/internal/uri"
)

func TestParse(t *testing.T) {
	type Test struct {
		Name    string
		Input   string
		Output  uri.URI
		IsValid bool
	}

	t.Parallel()

	const (
		standardURI  = "mongodb+srv://user:password@server.example.com"
		localhostURI = "mongodb://user:password@localhost:27017/db"
		badURI       = "https://example.com"
	)

	longURI := uri.Prefix + strings.Repeat("x", uri.MaxLength+1)

	tests := []Test{
		{"test standard uri", standardURI, uri.URI(standardURI), true},
		{"test localhost uri", localhostURI, uri.URI(localhostURI), true},
		{"test non-mongodb uri", badURI, "", false},
		{"test really long uri", longURI, "", false},
		{"test invalid uri", "", "", false},
	}

	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			t.Parallel()

			result, ok := uri.Parse(test.Input)
			if result != test.Output || ok != test.IsValid {
				t.Errorf(
					"uri.Parse(%v) = %v, %v; expected %v, %v",
					test.Input, result, ok, test.Output, test.IsValid)
			}
		})
	}
}
