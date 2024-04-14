// Copyright 2024 Aman Agnihotri
// SPDX-License-Identifier: AGPL-3.0-only

package lib

import (
	"context"
	"log"
	"os"
	"os/signal"
	"sync"
	"syscall"

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

func EnsureWatchers(base base.Wire) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	var watchers = []db.Watcher{
		base.AppStore(),
	}

	var waitGroup sync.WaitGroup

	done := make(chan os.Signal, 1)
	signal.Notify(done, syscall.SIGINT, syscall.SIGTERM)

	for _, watcher := range watchers {
		waitGroup.Add(1)

		go func(w db.Watcher) {
			defer waitGroup.Done()

			err := w.Watch(ctx)
			log.Println(err)
		}(watcher)
	}

	<-done
	cancel()
	waitGroup.Wait()
}
