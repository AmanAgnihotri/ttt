// Copyright 2024 Aman Agnihotri
// SPDX-License-Identifier: AGPL-3.0-only

package app

import (
	"context"
	"fmt"
	"sync"
	"sync/atomic"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"

	"ttt/internal/base/domain/app"
	"ttt/internal/base/infra/data/schema"
	"ttt/pkg/infra/data/db"
)

type Store struct {
	id         string
	collection *mongo.Collection
	app        atomic.Pointer[app.App]
	sync.Once
}

func NewStore(id string, db db.Service) *Store {
	return &Store{
		id:         id,
		collection: db.Config().Collection("app"),
		app:        atomic.Pointer[app.App]{},
		Once:       sync.Once{},
	}
}

func (s *Store) GetApp(ctx context.Context) *app.App {
	s.Once.Do(func() {
		s.app.Store(s.newApp(s.loadDao(ctx)))
	})

	return s.app.Load()
}

func (s *Store) Watch(ctx context.Context) error {
	err := db.Watch(ctx, s.collection, s.id, s.refresh)
	if err != nil {
		return fmt.Errorf("error during app watch: %w", err)
	}

	return nil
}

func (s *Store) refresh(dao DAO) {
	s.app.Store(s.newApp(dao))
}

func (s *Store) newApp(dao DAO) *app.App {
	switch dao.ID {
	case s.id:
		return app.NewApp(dao.MinimumVersion, dao.CurrentVersion)
	default:
		return app.NewDefaultApp()
	}
}

func (s *Store) loadDao(ctx context.Context) DAO {
	var dao DAO

	filter := bson.M{schema.ID: s.id}

	_ = s.collection.FindOne(ctx, filter).Decode(&dao)

	return dao
}
