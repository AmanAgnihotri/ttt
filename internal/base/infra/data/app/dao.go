// Copyright 2024 Aman Agnihotri
// SPDX-License-Identifier: AGPL-3.0-only

package app

import (
	"ttt/internal/base/domain/app"
)

type DAO struct {
	ID             string      `bson:"_id"`
	MinimumVersion app.Version `bson:"minimumVersion"`
	CurrentVersion app.Version `bson:"currentVersion"`
}
