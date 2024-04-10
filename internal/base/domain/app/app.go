// Copyright 2024 Aman Agnihotri
// SPDX-License-Identifier: AGPL-3.0-only

package app

type App struct {
	minimumVersion Version
	currentVersion Version
}

func NewDefaultApp() *App {
	return NewApp(1, 1)
}

func NewApp(minimumVersion, currentVersion Version) *App {
	return &App{minimumVersion, currentVersion}
}

func (a *App) IsValid(version Version) bool {
	return version >= a.minimumVersion
}

func (a *App) CanUpdate(version Version) bool {
	return version < a.currentVersion
}
