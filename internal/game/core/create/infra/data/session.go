// Copyright 2024 Aman Agnihotri
// SPDX-License-Identifier: AGPL-3.0-only

package data

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"golang.org/x/sync/errgroup"

	"ttt/internal/base/domain/game/user"
	"ttt/internal/base/infra/data/schema"
	"ttt/internal/game/core/create/domain/entity"
)

type session struct {
	context        context.Context
	session        mongo.Session
	userCollection *mongo.Collection
	gameCollection *mongo.Collection
}

func (s *session) Get(userID user.ID) (*entity.User, error) {
	filter := bson.M{schema.ID: userID}

	opts := options.FindOne().SetProjection(bson.M{
		schema.ID:         1,
		schema.Nick:       1,
		schema.State:      1,
		schema.AppVersion: 1,
		schema.Game:       1,
	})

	var dao entity.User

	err := s.userCollection.FindOne(s.context, filter, opts).Decode(&dao)
	if err != nil {
		return nil, fmt.Errorf("error finding user: %w", err)
	}

	return &dao, nil
}

func (s *session) Set(game *entity.Game, user *entity.User) error {
	group, ctx := errgroup.WithContext(s.context)

	group.Go(func() error {
		_, err := s.gameCollection.InsertOne(ctx, game, options.InsertOne())
		if err != nil {
			return fmt.Errorf("error inserting game: %w", err)
		}

		return nil
	})

	group.Go(func() error {
		update := bson.M{
			"$set": bson.M{
				schema.UpdateTime: user.UpdateTime,
				schema.Game:       user.Game,
			},
		}

		_, err := s.userCollection.UpdateByID(ctx, user.ID, update)
		if err != nil {
			return fmt.Errorf("error updating user: %w", err)
		}

		return nil
	})

	if err := group.Wait(); err != nil {
		if abortErr := s.session.AbortTransaction(s.context); abortErr != nil {
			return fmt.Errorf("error aborting transaction: %w", abortErr)
		}

		return fmt.Errorf("error updating documents: %w", err)
	}

	if err := s.session.CommitTransaction(s.context); err != nil {
		return fmt.Errorf("error committing transaction: %w", err)
	}

	return nil
}

func (s *session) End() {
	s.session.EndSession(s.context)
}
