// Copyright 2024 Aman Agnihotri
// SPDX-License-Identifier: AGPL-3.0-only

package user

import (
	"time"

	"ttt/internal/base/domain/app"
	"ttt/internal/base/domain/user/auth"
	"ttt/internal/base/domain/user/game"
	"ttt/internal/base/domain/user/guest"
)

type User struct {
	ID         ID           `bson:"_id"`
	Nick       Nick         `bson:"nick"`
	Group      Group        `bson:"group"`
	State      State        `bson:"state"`
	Platform   app.Platform `bson:"platform"`
	AppVersion app.Version  `bson:"appVersion"`
	CreateTime time.Time    `bson:"createTime"`
	UpdateTime time.Time    `bson:"updateTime"`
	ExpiryTime *time.Time   `bson:"expiryTime,omitempty"`
	Guest      guest.Guest  `bson:"guest"`
	Auth       *auth.Auth   `bson:"auth,omitempty"`
	Game       *game.Game   `bson:"game,omitempty"`
}
