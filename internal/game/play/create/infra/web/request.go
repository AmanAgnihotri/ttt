// Copyright 2024 Aman Agnihotri
// SPDX-License-Identifier: AGPL-3.0-only

package web

import (
	"ttt/internal/base/domain/game/user"
)

type RequestType string

const (
	PlayRequest RequestType = "play"
)

type request struct {
	UserID user.ID        `json:"-"`
	Type   RequestType    `json:"type"`
	Data   map[string]any `json:"data"`
}

type Play struct {
	Position int `json:"position"`
}
