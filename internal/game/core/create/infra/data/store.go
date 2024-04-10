// Copyright 2024 Aman Agnihotri
// SPDX-License-Identifier: AGPL-3.0-only

package data

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"ttt/internal/base/domain/app"
	"ttt/internal/base/domain/game"
	"ttt/internal/base/infra/data/schema"
	"ttt/internal/game/core/create/domain"
	"ttt/pkg/infra/data/db"
)

func NewStore(db db.Service, appStore app.Store) domain.Store {
	return &store{
		db.Client(),
		db.Core().Collection("user"),
		db.Core().Collection("game"),
		appStore,
	}
}

type store struct {
	client         *mongo.Client
	userCollection *mongo.Collection
	gameCollection *mongo.Collection
	app.Store
}

func (s *store) IsIDAvailable(ctx context.Context, gameID game.ID) bool {
	filter := bson.M{schema.ID: gameID}

	return s.isAvailable(ctx, filter)
}

func (s *store) IsCodeAvailable(ctx context.Context, gameCode game.Code) bool {
	filter := bson.M{schema.Code: gameCode}

	return s.isAvailable(ctx, filter)
}

func (s *store) NewSession(ctx context.Context) domain.Session {
	newSession, err := s.client.StartSession()
	if err != nil {
		return nil
	}

	if err = newSession.StartTransaction(); err != nil {
		return nil
	}

	return &session{ctx, newSession, s.userCollection, s.gameCollection}
}

func (s *store) isAvailable(ctx context.Context, filter bson.M) bool {
	opts := options.Count().SetLimit(1)

	count, err := s.gameCollection.CountDocuments(ctx, filter, opts)

	return err == nil && count == 0
}
