// Copyright 2024 Aman Agnihotri
// SPDX-License-Identifier: AGPL-3.0-only

package game

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
		collection: db.Core().Collection("game"),
	}
}

func (s *Store) EnsureIndexes(ctx context.Context) error {
	codeIndex := mongo.IndexModel{
		Keys:    bson.D{{Key: "code", Value: 1}},
		Options: options.Index().SetUnique(true).SetSparse(true),
	}

	expiryTimeIndex := mongo.IndexModel{
		Keys:    bson.D{{Key: "expiryTime", Value: 1}},
		Options: options.Index().SetExpireAfterSeconds(0),
	}

	indexModels := []mongo.IndexModel{codeIndex, expiryTimeIndex}

	_, err := s.collection.Indexes().CreateMany(ctx, indexModels)
	if err != nil {
		return fmt.Errorf("error ensuring indexes: %w", err)
	}

	return nil
}
