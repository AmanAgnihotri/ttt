// Copyright 2024 Aman Agnihotri
// SPDX-License-Identifier: AGPL-3.0-only

package data

import (
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"golang.org/x/sync/errgroup"

	"ttt/internal/base/domain/game"
	"ttt/internal/base/domain/game/user"
	"ttt/internal/base/infra/data/schema"
	"ttt/internal/game/play/create/domain"
	"ttt/pkg/infra/data/db"
)

func NewStore(db db.Service) domain.Store {
	return &store{
		db.Core().Collection("user"),
		db.Core().Collection("game"),
	}
}

type store struct {
	userCollection *mongo.Collection
	gameCollection *mongo.Collection
}

func (s *store) IsActive(gameID game.ID) bool {
	const timeout = 3 * time.Second

	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	filter := bson.M{
		schema.ID: gameID,
		schema.FinishTime: bson.M{
			"$exists": false,
		},
	}

	opts := options.Count().SetLimit(1)

	count, err := s.gameCollection.CountDocuments(ctx, filter, opts)

	return err == nil && count == 1
}

func (s *store) Get(gameID game.ID) (*domain.Game, error) {
	const timeout = 3 * time.Second

	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	filter := bson.M{schema.ID: gameID}

	opts := options.FindOne().SetProjection(bson.M{
		schema.ID:    1,
		schema.Users: 1,
	})

	var dao domain.Game

	err := s.gameCollection.FindOne(ctx, filter, opts).Decode(&dao)
	if err != nil {
		return nil, fmt.Errorf("error finding game: %w", err)
	}

	return &dao, nil
}

func (s *store) Set(game *domain.Game) error {
	const timeout = 3 * time.Second

	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	group, ctx := errgroup.WithContext(ctx)

	unset := func(userID user.ID) {
		group.Go(func() error {
			update := bson.M{
				"$unset": bson.M{
					schema.Game: 1,
				},
			}

			_, err := s.userCollection.UpdateByID(ctx, userID, update)
			if err != nil {
				return fmt.Errorf("error updating user: %w", err)
			}

			return nil
		})
	}

	unset(game.Users[0].ID)
	unset(game.Users[1].ID)

	group.Go(func() error {
		update := bson.M{
			"$set": bson.M{
				schema.FinishTime: game.FinishTime,
				schema.Outcome:    game.Outcome,
				schema.Winner:     game.Winner,
				schema.Moves:      game.Moves,
			},
			"$unset": bson.M{
				schema.Code:       1,
				schema.ExpiryTime: 1,
			},
		}

		_, err := s.gameCollection.UpdateByID(ctx, game.ID, update)
		if err != nil {
			return fmt.Errorf("error updating game: %w", err)
		}

		return nil
	})

	if err := group.Wait(); err != nil {
		return fmt.Errorf("error updating documents: %w", err)
	}

	return nil
}
