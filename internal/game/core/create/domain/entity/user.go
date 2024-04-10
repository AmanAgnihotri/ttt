// Copyright 2024 Aman Agnihotri
// SPDX-License-Identifier: AGPL-3.0-only

package entity

import (
	"time"

	"ttt/internal/base/domain/app"
	"ttt/internal/base/domain/user"
	"ttt/internal/base/domain/user/game"
)

type User struct {
	ID         user.ID     `bson:"_id"`
	Nick       user.Nick   `bson:"nick"`
	State      user.State  `bson:"state"`
	AppVersion app.Version `bson:"appVersion"`
	UpdateTime time.Time   `bson:"updateTime"`
	Game       *game.Game  `bson:"game"`
}

func (u *User) CanCreateGame(now time.Time) bool {
	switch u.State {
	case user.Deleted, user.Banned, user.Locked:
		return false
	default:
		return u.Game == nil || u.Game.ExpiryTime.Before(now)
	}
}

func (u *User) SetGame(game *Game) {
	u.UpdateTime = game.CreateTime
	u.Game = newGame(game)
}

func newGame(g *Game) *game.Game {
	return &game.Game{
		Code:       game.Code(g.Code),
		ExpiryTime: g.ExpiryTime,
	}
}
