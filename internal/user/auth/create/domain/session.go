// Copyright 2024 Aman Agnihotri
// SPDX-License-Identifier: AGPL-3.0-only

package domain

import (
	"ttt/internal/base/domain/user"
	"ttt/internal/base/domain/user/auth"
)

type Session interface {
	Get(userID user.ID, authToken auth.Token) (*User, error)
	Set(user *User) error
	End()
}
