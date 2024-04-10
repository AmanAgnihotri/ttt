// Copyright 2024 Aman Agnihotri
// SPDX-License-Identifier: AGPL-3.0-only

package domain

import (
	"ttt/internal/base/domain/user"
)

type Session interface {
	Get(userID user.ID) (*User, error)
	Set(user *User) error
	End()
}
