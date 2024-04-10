// Copyright 2024 Aman Agnihotri
// SPDX-License-Identifier: AGPL-3.0-only

package domain

import (
	"time"

	"ttt/internal/base/domain/app"
	"ttt/internal/base/domain/user"
	"ttt/internal/base/domain/user/auth"
)

type User struct {
	ID         user.ID      `bson:"_id"`
	State      user.State   `bson:"state"`
	Platform   app.Platform `bson:"platform"`
	AppVersion app.Version  `bson:"appVersion"`
	UpdateTime time.Time    `bson:"updateTime"`
	Auth       *auth.Auth   `bson:"auth"`
}

func (u *User) Login(
	platform app.Platform,
	appVersion app.Version,
	now time.Time,
) Event {
	if event := u.checkState(); event != none {
		return event
	}

	newAuth, ok := auth.NewAuth(now)
	if !ok {
		return Error
	}

	u.Auth = newAuth
	u.Platform = platform
	u.AppVersion = appVersion
	u.UpdateTime = now

	return LoggedIn
}

func (u *User) checkState() Event {
	switch u.State {
	case user.Deleted:
		return Deleted
	case user.Banned:
		return Banned
	case user.Locked:
		return Locked
	default:
		return none
	}
}
