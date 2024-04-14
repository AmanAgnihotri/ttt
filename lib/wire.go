// Copyright 2024 Aman Agnihotri
// SPDX-License-Identifier: AGPL-3.0-only

package lib

import (
	"context"
	"log"

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

	ensureIndexes(baseWire)
}

func ensureIndexes(base base.Wire) {
	ctx := context.Background()

	var indexers = []db.Indexer{
		base.GameStore(),
		base.UserStore(),
	}

	for _, indexer := range indexers {
		err := indexer.EnsureIndexes(ctx)
		if err != nil {
			log.Println(err)
		}
	}
}
