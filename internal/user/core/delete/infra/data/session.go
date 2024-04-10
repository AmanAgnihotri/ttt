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
	"ttt/internal/user/core/delete/domain"
)

type session struct {
	context    context.Context
	session    mongo.Session
	collection *mongo.Collection
}

func (s *session) Get(userID user.ID) (*domain.User, error) {
	filter := bson.M{schema.ID: userID}

	opts := options.FindOne().SetProjection(bson.M{
		schema.ID:    1,
		schema.State: 1,
	})

	var dao domain.User

	err := s.collection.FindOne(s.context, filter, opts).Decode(&dao)
	if err != nil {
		return nil, fmt.Errorf("error finding user: %w", err)
	}

	return &dao, nil
}

func (s *session) Set(user *domain.User) error {
	update := bson.M{
		"$set": bson.M{
			schema.State:      user.State,
			schema.UpdateTime: user.UpdateTime,
			schema.ExpiryTime: user.ExpiryTime,
		},
	}

	if _, err := s.collection.UpdateByID(s.context, user.ID, update); err != nil {
		if abortErr := s.session.AbortTransaction(s.context); abortErr != nil {
			return fmt.Errorf("error aborting transaction: %w", abortErr)
		}

		return fmt.Errorf("error updating document: %w", err)
	}

	if err := s.session.CommitTransaction(s.context); err != nil {
		return fmt.Errorf("error committing transaction: %w", err)
	}

	return nil
}

func (s *session) End() {
	s.session.EndSession(s.context)
}
