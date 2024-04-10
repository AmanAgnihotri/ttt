// Copyright 2024 Aman Agnihotri
// SPDX-License-Identifier: AGPL-3.0-only

package game

import "time"

type Code string

type Game struct {
	Code       Code      `bson:"code"       json:"code"`
	ExpiryTime time.Time `bson:"expiryTime" json:"expiryTime"`
}
