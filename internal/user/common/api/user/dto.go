// Copyright 2024 Aman Agnihotri
// SPDX-License-Identifier: AGPL-3.0-only

package user

import (
	"time"

	"ttt/internal/base/domain/user"
	"ttt/internal/base/domain/user/game"
	"ttt/internal/user/common/api/user/login"
)

type DTO struct {
	ID           user.ID    `json:"id"`
	Nick         user.Nick  `json:"nick"`
	Group        user.Group `json:"group"`
	State        user.State `json:"state"`
	CreateTime   time.Time  `json:"createTime"`
	UpdateTime   time.Time  `json:"updateTime"`
	ExpiryTime   *time.Time `json:"expiryTime,omitempty"`
	LoginID      login.ID   `json:"loginId"`
	CanUpdateApp bool       `json:"canUpdateApp,omitempty"`
	Game         *game.Game `json:"game,omitempty"`
}
