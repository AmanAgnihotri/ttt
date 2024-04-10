// Copyright 2024 Aman Agnihotri
// SPDX-License-Identifier: AGPL-3.0-only

package uri

import "strings"

const (
	MaxLength = 300
	Prefix    = "mongodb"
)

type URI string

func Parse(s string) (URI, bool) {
	if isValid(s) {
		return URI(s), true
	}

	return "", false
}

func isValid(s string) bool {
	return len(s) <= MaxLength && strings.HasPrefix(s, Prefix)
}

func (u URI) String() string {
	return string(u)
}
