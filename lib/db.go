// Copyright 2024 Aman Agnihotri
// SPDX-License-Identifier: AGPL-3.0-only

package lib

import (
	"context"
	"fmt"

	"ttt/pkg/infra/data/db"
)

const (
	connectionStringKey = "MongoDB_ConnectionString"
	databasePrefixKey   = "MongoDB_DatabasePrefix"
)

func NewDBContext(ctx context.Context) (*db.Context, error) {
	config, err := newDBConfig()
	if err != nil {
		return nil, fmt.Errorf("error creating db config: %w", err)
	}

	dbContext, err := db.NewContext(ctx, config)
	if err != nil {
		return nil, fmt.Errorf("error creating db context: %w", err)
	}

	return dbContext, nil
}

func newDBConfig() (db.Config, error) {
	var config db.Config

	connectionString, err := LookupEnv(connectionStringKey)
	if err != nil {
		return config, err
	}

	databasePrefix, err := LookupEnv(databasePrefixKey)
	if err != nil {
		return config, err
	}

	config, err = db.NewConfig(connectionString, databasePrefix)
	if err != nil {
		return config, fmt.Errorf("error creating db config: %w", err)
	}

	return config, nil
}
