// Copyright 2024 Aman Agnihotri
// SPDX-License-Identifier: AGPL-3.0-only

package data

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"ttt/internal/base/domain/app"
	"ttt/internal/base/domain/user"
	"ttt/internal/base/infra/data/schema"
	"ttt/internal/user/core/create/domain"
	"ttt/pkg/infra/data/db"
)

func NewStore(db db.Service, appStore app.Store) domain.Store {
	return &store{
		db.Core().Collection("user"),
		appStore,
	}
}

type store struct {
	collection *mongo.Collection
	app.Store
}

func (s *store) IsIDAvailable(ctx context.Context, userID user.ID) bool {
	filter := bson.M{schema.ID: userID}

	return s.isAvailable(ctx, filter)
}

func (s *store) IsNickAvailable(ctx context.Context, nick user.Nick) bool {
	filter := bson.M{schema.Nick: nick}

	return s.isAvailable(ctx, filter)
}

func (s *store) Insert(ctx context.Context, user *domain.User) bool {
	opts := options.InsertOne()

	_, err := s.collection.InsertOne(ctx, user, opts)

	return err == nil
}

func (s *store) isAvailable(ctx context.Context, filter bson.M) bool {
	opts := options.Count().SetLimit(1)

	count, err := s.collection.CountDocuments(ctx, filter, opts)

	return err == nil && count == 0
}
