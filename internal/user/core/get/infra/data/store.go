// Copyright 2024 Aman Agnihotri
// SPDX-License-Identifier: AGPL-3.0-only

package data

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"ttt/internal/base/domain/user"
	"ttt/internal/base/infra/data/schema"
	"ttt/internal/user/core/get/domain"
	"ttt/pkg/infra/data/db"
)

func NewStore(db db.Service) domain.Store {
	return &store{
		db.Core().Collection("user"),
	}
}

type store struct {
	collection *mongo.Collection
}

func (s *store) Get(ctx context.Context, userID user.ID) (*user.User, error) {
	filter := bson.M{schema.ID: userID}

	var dao user.User

	err := s.collection.FindOne(ctx, filter, options.FindOne()).Decode(&dao)
	if err != nil {
		return nil, fmt.Errorf("error finding user: %w", err)
	}

	return &dao, nil
}
