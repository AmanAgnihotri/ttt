// Copyright 2024 Aman Agnihotri
// SPDX-License-Identifier: AGPL-3.0-only

package domain

import (
	"time"

	"ttt/internal/base/domain/user"
)

type User struct {
	ID         user.ID    `bson:"_id"`
	State      user.State `bson:"state"`
	UpdateTime time.Time  `bson:"updateTime"`
	ExpiryTime time.Time  `bson:"expiryTime"`
}

func (u *User) Delete(now time.Time) Event {
	if event := u.checkState(); event != none {
		return event
	}

	u.State = user.Deleted

	u.UpdateTime = now
	u.ExpiryTime = now.Add(7 * 24 * time.Hour)

	return Deleted
}

func (u *User) checkState() Event {
	switch u.State {
	case user.Deleted:
		return AlreadyDeleted
	case user.Banned:
		return Banned
	case user.Locked:
		return Locked
	default:
		return none
	}
}
