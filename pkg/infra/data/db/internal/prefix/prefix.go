// Copyright 2024 Aman Agnihotri
// SPDX-License-Identifier: AGPL-3.0-only

package prefix

import "strings"

const (
	MinLength = 1
	MaxLength = 12
)

type Prefix string

func Parse(s string) (Prefix, bool) {
	if isValid(s) {
		return Prefix(s), true
	}

	return "", false
}

func isValid(s string) bool {
	return len(s) >= MinLength && len(s) <= MaxLength &&
		!strings.ContainsAny(s, ".%$*<>:|?%^#@!();[]{}, ")
}

func (p Prefix) String() string {
	return string(p)
}
