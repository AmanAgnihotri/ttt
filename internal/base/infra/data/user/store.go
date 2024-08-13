// Copyright 2024 Aman Agnihotri
// SPDX-License-Identifier: AGPL-3.0-only

package user

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"ttt/pkg/infra/data/db"
)

type Store struct {
	collection *mongo.Collection
}

func NewStore(db db.Service) *Store {
	return &Store{
		collection: db.Core().Collection("user"),
	}
}

func (s *Store) EnsureIndexes(ctx context.Context) error {
	nickIndex := mongo.IndexModel{
		Keys:    bson.D{{Key: "nick", Value: 1}},
		Options: options.Index().SetUnique(true),
	}

	expiryTimeIndex := mongo.IndexModel{
		Keys:    bson.D{{Key: "expiryTime", Value: 1}},
		Options: options.Index().SetExpireAfterSeconds(0),
	}

	indexModels := []mongo.IndexModel{nickIndex, expiryTimeIndex}

	_, err := s.collection.Indexes().CreateMany(ctx, indexModels)
	if err != nil {
		return fmt.Errorf("error ensuring indexes: %w", err)
	}

	return nil
}
