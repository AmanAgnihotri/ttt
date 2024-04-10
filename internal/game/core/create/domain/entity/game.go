// Copyright 2024 Aman Agnihotri
// SPDX-License-Identifier: AGPL-3.0-only

package entity

import (
	"time"

	"ttt/internal/base/domain/game"
	"ttt/internal/base/domain/game/user"
)

const initialLifetime = 30 * time.Minute

type Game struct {
	ID         game.ID     `bson:"_id"`
	Code       game.Code   `bson:"code"`
	CreateTime time.Time   `bson:"createTime"`
	ExpiryTime time.Time   `bson:"expiryTime"`
	Users      []user.User `bson:"users"`
}

func NewGame(
	gameID game.ID,
	gameCode game.Code,
	createTime time.Time,
	user *User,
	side user.Side,
) *Game {
	return &Game{
		ID:         gameID,
		Code:       gameCode,
		CreateTime: createTime,
		ExpiryTime: createTime.Add(initialLifetime),
		Users:      newUsers(user, side),
	}
}

func newUsers(u *User, side user.Side) []user.User {
	return []user.User{newUser(u, side)}
}

func newUser(u *User, side user.Side) user.User {
	return user.User{
		ID:   user.ID(u.ID),
		Nick: user.Nick(u.Nick),
		Side: side.Resolve(),
	}
}
