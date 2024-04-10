// Copyright 2024 Aman Agnihotri
// SPDX-License-Identifier: AGPL-3.0-only

package db

import (
	"errors"

	"ttt/pkg/infra/data/db/internal/prefix"
	"ttt/pkg/infra/data/db/internal/uri"
)

type Config struct {
	uri    uri.URI
	prefix prefix.Prefix
}

var (
	ErrInvalidURI    = errors.New("invalid mongodb uri")
	ErrInvalidPrefix = errors.New("invalid database prefix")
)

func NewConfig(connectionString, databasePrefix string) (Config, error) {
	var config Config

	dbURI, ok := uri.Parse(connectionString)
	if !ok {
		return config, ErrInvalidURI
	}

	dbPrefix, ok := prefix.Parse(databasePrefix)
	if !ok {
		return config, ErrInvalidPrefix
	}

	config = Config{dbURI, dbPrefix}

	return config, nil
}
