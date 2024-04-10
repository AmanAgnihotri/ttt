// Copyright 2024 Aman Agnihotri
// SPDX-License-Identifier: AGPL-3.0-only

package base

import (
	"ttt/internal/base/infra/data/app"
	"ttt/pkg"
)

type Wire interface {
	pkg.Wire
	AppStore() *app.Store
}

func NewWire(appID string, pkg pkg.Wire) Wire {
	return &wire{
		pkg,
		app.NewStore(appID, pkg.DBService()),
	}
}

type wire struct {
	pkg.Wire
	appStore *app.Store
}

func (w *wire) AppStore() *app.Store { return w.appStore }
