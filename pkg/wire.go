// Copyright 2024 Aman Agnihotri
// SPDX-License-Identifier: AGPL-3.0-only

package pkg

import (
	"ttt/pkg/api/jwt"
	"ttt/pkg/infra/data/db"
	"ttt/pkg/infra/web/http"
	"ttt/pkg/infra/web/time"
)

type Wire interface {
	Timer() time.Timer
	DBService() db.Service
	JWTService() jwt.Service
	Router() http.Router
}

func NewWire(db db.Service, jwt jwt.Service, router http.Router) Wire {
	return &wire{
		time.NewTimer(),
		db,
		jwt,
		router,
	}
}

type wire struct {
	timer      time.Timer
	dbService  db.Service
	jwtService jwt.Service
	router     http.Router
}

func (w *wire) Timer() time.Timer       { return w.timer }
func (w *wire) DBService() db.Service   { return w.dbService }
func (w *wire) JWTService() jwt.Service { return w.jwtService }
func (w *wire) Router() http.Router     { return w.router }
