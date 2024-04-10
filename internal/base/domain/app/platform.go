// Copyright 2024 Aman Agnihotri
// SPDX-License-Identifier: AGPL-3.0-only

package app

const (
	android = "android"
	ios     = "ios"
	editor  = "editor"
	browser = "browser"
)

type Platform string

const (
	Android Platform = android
	IOS     Platform = ios
	Editor  Platform = editor
	Browser Platform = browser
)

func ParsePlatform(s string) (Platform, bool) {
	switch s {
	case android, ios, editor, browser:
		return Platform(s), true
	default:
		return "", false
	}
}
