// Copyright 2024 Aman Agnihotri
// SPDX-License-Identifier: AGPL-3.0-only

package entity

import (
	"time"

	"ttt/internal/base/domain/game"
	"ttt/internal/base/domain/game/user"
)

type Game struct {
	ID         game.ID     `bson:"_id"`
	Code       game.Code   `bson:"code"`
	CreateTime time.Time   `bson:"createTime"`
	UpdateTime time.Time   `bson:"updateTime"`
	ExpiryTime time.Time   `bson:"expiryTime"`
	Users      []user.User `bson:"users"`
}

func (g *Game) IsFull() bool {
	const maxUserCount = 2

	return len(g.Users) == maxUserCount
}

func (g *Game) HasUser(u *User) bool {
	userID := user.ID(u.ID)

	return g.hasUserID(userID)
}

func (g *Game) AddUser(user *User, time time.Time) {
	g.UpdateTime = time

	oppositeSide := g.Users[0].Side.Opposite()

	g.Users = append(g.Users, newUser(user, oppositeSide))
}

func (g *Game) ExtendExpiryTime() {
	const extraLifetime = 10 * time.Minute

	g.ExpiryTime = g.ExpiryTime.Add(extraLifetime)
}

func (g *Game) hasUserID(userID user.ID) bool {
	for _, u := range g.Users {
		if u.ID == userID {
			return true
		}
	}

	return false
}

func newUser(u *User, side user.Side) user.User {
	return user.User{
		ID:   user.ID(u.ID),
		Nick: user.Nick(u.Nick),
		Side: side,
	}
}
