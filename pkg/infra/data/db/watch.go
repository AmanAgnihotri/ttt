// Copyright 2024 Aman Agnihotri
// SPDX-License-Identifier: AGPL-3.0-only

package db

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func Watch[T any](
	ctx context.Context,
	collection *mongo.Collection,
	documentID any,
	apply func(T),
) error {
	opts := options.ChangeStream().SetFullDocument(options.UpdateLookup)

	pipeline := mongo.Pipeline{bson.D{{
		Key: "$match", Value: bson.M{
			"documentKey._id": documentID,
			"operationType": bson.M{
				"$in": bson.A{"insert", "update", "replace", "delete"},
			},
		},
	}}}

	changeStream, err := collection.Watch(ctx, pipeline, opts)
	if err != nil {
		return fmt.Errorf("error opening change stream: %w", err)
	}

	defer func() {
		closeErr := changeStream.Close(ctx)
		if closeErr != nil && err == nil {
			err = fmt.Errorf("error closing change stream: %w", closeErr)
		}
	}()

	if err = process(ctx, changeStream, apply); err != nil {
		return fmt.Errorf("error processing change stream: %w", err)
	}

	return err
}

func process[T any](
	ctx context.Context,
	changeStream *mongo.ChangeStream,
	apply func(T),
) error {
	for changeStream.Next(ctx) {
		var changeEvent bson.M
		if err := changeStream.Decode(&changeEvent); err != nil {
			return fmt.Errorf("error decoding change event: %w", err)
		}

		document, err := getDocument[T](changeEvent)
		if err != nil {
			return fmt.Errorf("error getting document from change event: %w", err)
		}

		apply(document)
	}

	if err := changeStream.Err(); err != nil {
		return fmt.Errorf("error in change stream: %w", err)
	}

	return nil
}

func getDocument[T any](change bson.M) (T, error) {
	var document T

	if fullDocument, ok := change["fullDocument"].(bson.M); ok {
		rawData, err := bson.Marshal(fullDocument)
		if err != nil {
			return document, fmt.Errorf("error marshalling document: %w", err)
		}

		err = bson.Unmarshal(rawData, &document)
		if err != nil {
			return document, fmt.Errorf("error unmarshalling document: %w", err)
		}
	}

	return document, nil
}
