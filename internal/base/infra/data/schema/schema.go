// Copyright 2024 Aman Agnihotri
// SPDX-License-Identifier: AGPL-3.0-only

package schema

const (
	ID = "_id"

	CreateTime = "createTime"
	UpdateTime = "updateTime"
	FinishTime = "finishTime"
	ExpiryTime = "expiryTime"

	Code    = "code"
	Outcome = "outcome"

	Winner       = "winner"
	WinnerID     = "winner.id"
	WinnerMethod = "winner.method"

	Users    = "users"
	UserID   = "user.id"
	UserNick = "user.nick"
	UserSide = "user.side"

	Moves = "moves"

	Nick       = "nick"
	Group      = "group"
	State      = "state"
	Platform   = "platform"
	AppVersion = "appVersion"

	Guest      = "guest"
	GuestToken = "guest.token"

	Auth           = "auth"
	AuthToken      = "auth.token"
	AuthExpiryTime = "auth.expiryTime"

	Game           = "game"
	GameID         = "game.id"
	GameExpiryTime = "game.expiryTime"
)
