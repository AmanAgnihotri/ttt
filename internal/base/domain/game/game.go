// Copyright 2024 Aman Agnihotri
// SPDX-License-Identifier: AGPL-3.0-only

package game

import (
	"time"

	"ttt/internal/base/domain/game/user"
	"ttt/internal/base/domain/game/winner"
)

type Game struct {
	ID         ID             `bson:"_id"`
	Code       Code           `bson:"code,omitempty"`
	CreateTime time.Time      `bson:"createTime"`
	UpdateTime *time.Time     `bson:"updateTime,omitempty"`
	FinishTime *time.Time     `bson:"finishTime,omitempty"`
	ExpiryTime *time.Time     `bson:"expiryTime,omitempty"`
	Outcome    *Outcome       `bson:"outcome,omitempty"`
	Winner     *winner.Winner `bson:"winner,omitempty"`
	Users      []user.User    `bson:"users"`
	Moves      []Move         `bson:"moves,omitempty"`
}
