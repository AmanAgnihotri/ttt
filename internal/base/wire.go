// Copyright 2024 Aman Agnihotri
// SPDX-License-Identifier: AGPL-3.0-only

package base

import (
	"ttt/internal/base/infra/data/app"
	"ttt/internal/base/infra/data/game"
	"ttt/internal/base/infra/data/user"
	"ttt/pkg"
)

type Wire interface {
	pkg.Wire
	AppStore() *app.Store
	GameStore() *game.Store
	UserStore() *user.Store
}

func NewWire(appID string, pkg pkg.Wire) Wire {
	return &wire{
		pkg,
		app.NewStore(appID, pkg.DBService()),
		game.NewStore(pkg.DBService()),
		user.NewStore(pkg.DBService()),
	}
}

type wire struct {
	pkg.Wire
	appStore  *app.Store
	gameStore *game.Store
	userStore *user.Store
}

func (w *wire) AppStore() *app.Store   { return w.appStore }
func (w *wire) GameStore() *game.Store { return w.gameStore }
func (w *wire) UserStore() *user.Store { return w.userStore }
