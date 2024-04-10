// Copyright 2024 Aman Agnihotri
// SPDX-License-Identifier: AGPL-3.0-only

package db

import (
	"go.mongodb.org/mongo-driver/mongo"
)

type Service interface {
	Client() *mongo.Client
	Config() *mongo.Database
	Core() *mongo.Database
}
