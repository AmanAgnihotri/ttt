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
	"ttt/internal/base/domain/user/guest"
	"ttt/internal/base/infra/data/schema"
	"ttt/internal/user/login/create/domain"
)

type session struct {
	context    context.Context
	session    mongo.Session
	collection *mongo.Collection
}

func (s *session) Get(
	userID user.ID,
	guestToken guest.Token,
) (*domain.User, error) {
	filter := bson.M{
		schema.ID:         userID,
		schema.GuestToken: guestToken,
	}

	opts := options.FindOne().SetProjection(bson.M{
		schema.ID:    1,
		schema.State: 1,
		schema.Auth:  1,
	})

	var dao domain.User

	err := s.collection.FindOne(s.context, filter, opts).Decode(&dao)
	if err != nil {
		return nil, fmt.Errorf("error finding user: %w", err)
	}

	return &dao, nil
}

func (s *session) Set(user *domain.User) (*user.User, error) {
	filter := bson.M{schema.ID: user.ID}

	update := bson.M{
		"$set": bson.M{
			schema.Platform:   user.Platform,
			schema.AppVersion: user.AppVersion,
			schema.UpdateTime: user.UpdateTime,
			schema.Auth:       user.Auth,
		},
	}

	opts := options.FindOneAndUpdate().SetReturnDocument(options.After)

	result := s.collection.FindOneAndUpdate(s.context, filter, update, opts)

	var dao domain.UserData
	if err := result.Decode(&dao); err != nil {
		if abortErr := s.session.AbortTransaction(s.context); abortErr != nil {
			return nil, fmt.Errorf("error aborting transaction: %w", abortErr)
		}

		return nil, fmt.Errorf("error decoding user: %w", err)
	}

	if err := s.session.CommitTransaction(s.context); err != nil {
		return nil, fmt.Errorf("error committing transaction: %w", err)
	}

	return &dao, nil
}

func (s *session) End() {
	s.session.EndSession(s.context)
}
