// Copyright 2024 Aman Agnihotri
// SPDX-License-Identifier: AGPL-3.0-only

package db

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

type Context struct {
	client *mongo.Client
	config *mongo.Database
	core   *mongo.Database
}

func NewContext(ctx context.Context, config Config) (*Context, error) {
	opts := options.Client().ApplyURI(config.uri.String())

	client, err := mongo.Connect(ctx, opts)
	if err != nil {
		return nil, fmt.Errorf("error during connect: %w", err)
	}

	return newContext(client, config.prefix.String()), nil
}

func newContext(client *mongo.Client, prefix string) *Context {
	return &Context{
		client: client,
		config: client.Database(prefix + "_config"),
		core:   client.Database(prefix + "_core"),
	}
}

func (c *Context) Ping(ctx context.Context) error {
	err := c.client.Ping(ctx, readpref.Primary())
	if err != nil {
		return fmt.Errorf("error during ping: %w", err)
	}

	return nil
}

func (c *Context) Disconnect(ctx context.Context) error {
	err := c.client.Disconnect(ctx)
	if err != nil {
		return fmt.Errorf("error during disconnect: %w", err)
	}

	return nil
}

func (c *Context) Client() *mongo.Client   { return c.client }
func (c *Context) Config() *mongo.Database { return c.config }
func (c *Context) Core() *mongo.Database   { return c.core }
