// Copyright 2024 Aman Agnihotri
// SPDX-License-Identifier: AGPL-3.0-only

package lib

import (
	"ttt/internal/base"
	"ttt/internal/game"
	"ttt/internal/user"
	"ttt/pkg"
	"ttt/pkg/api/jwt"
	"ttt/pkg/infra/data/db"
	"ttt/pkg/infra/web/http"
)

func Configure(appID string, db *db.Context, jwt *jwt.Context, mux *http.Mux) {
	pkgWire := pkg.NewWire(db, jwt, mux)
	baseWire := base.NewWire(appID, pkgWire)

	game.Configure(baseWire)
	user.Configure(baseWire)
}
