// Copyright 2024 Aman Agnihotri
// SPDX-License-Identifier: AGPL-3.0-only

package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"ttt/lib"
	"ttt/pkg/infra/data/db"
	"ttt/pkg/infra/web/http"
)

const (
	name    = "ttt"
	address = ":8080"
	maxLoad = 20_000
)

func main() {
	const timeout = 3 * time.Second

	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	dbContext, err := lib.NewDBContext(ctx)
	{
		if err != nil {
			log.Println(err)

			return
		}

		if err = dbContext.Ping(ctx); err != nil {
			log.Println(err)

			return
		}

		log.Println("connected with mongodb;")
	}

	jwtContext, err := lib.NewJWTContext()
	if err != nil {
		log.Println(err)

		return
	}

	mux := http.NewMux(jwtContext)

	mux.Handle("/", func(ctx *http.Context) {
		ctx.PlainText("ok")
	})

	lib.Configure(name, dbContext, jwtContext, mux)

	server := http.NewServer(name, address, maxLoad, mux)

	if err = server.Start(); err != nil {
		log.Println(err)

		return
	}

	log.Printf("%s listening on %s\n", name, address)

	done := make(chan os.Signal, 1)
	signal.Notify(done, syscall.SIGINT, syscall.SIGTERM)
	<-done

	shutdown(server, dbContext)
}

func shutdown(server *http.Server, dbClient db.Client) {
	const timeout = 5 * time.Second

	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	log.Println("shutdown initiated;")

	if err := server.Stop(ctx); err != nil {
		log.Println(err)
	} else {
		log.Println("server shutdown gracefully;")
	}

	if err := dbClient.Disconnect(ctx); err != nil {
		log.Println(err)
	} else {
		log.Println("mongodb disconnected gracefully;")
	}
}
