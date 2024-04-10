// Copyright 2024 Aman Agnihotri
// SPDX-License-Identifier: AGPL-3.0-only

package app

type Version int32

func ParseVersion(v int32) (Version, bool) {
	if v > 0 {
		return Version(v), true
	}

	return 0, false
}
