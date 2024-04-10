// Copyright 2024 Aman Agnihotri
// SPDX-License-Identifier: AGPL-3.0-only

package game

import (
	"time"

	"ttt/internal/base/domain/game"
)

type DTO struct {
	Code       game.Code `json:"code"`
	CreateTime time.Time `json:"createTime"`
	ExpiryTime time.Time `json:"expiryTime"`
}
