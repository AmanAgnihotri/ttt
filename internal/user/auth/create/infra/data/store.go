// Copyright 2024 Aman Agnihotri
// SPDX-License-Identifier: AGPL-3.0-only

package data

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"

	"ttt/internal/base/domain/app"
	"ttt/internal/user/auth/create/domain"
	"ttt/pkg/infra/data/db"
)

func NewStore(db db.Service, appStore app.Store) domain.Store {
	return &store{
		db.Client(),
		db.Core().Collection("user"),
		appStore,
	}
}

type store struct {
	client     *mongo.Client
	collection *mongo.Collection
	app.Store
}

func (s *store) NewSession(ctx context.Context) domain.Session {
	newSession, err := s.client.StartSession()
	if err != nil {
		return nil
	}

	if err = newSession.StartTransaction(); err != nil {
		return nil
	}

	return &session{ctx, newSession, s.collection}
}
