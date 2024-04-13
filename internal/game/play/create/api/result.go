// Copyright 2024 Aman Agnihotri
// SPDX-License-Identifier: AGPL-3.0-only

package api

import (
	"ttt/internal/base/domain/game"
	"ttt/internal/base/domain/game/user"
	"ttt/internal/base/domain/game/winner"
)

type Result interface{}

type StartedResult struct {
	ID         game.ID     `json:"id"`
	Users      []user.User `json:"users"`
	State      string      `json:"state"`
	Marks      string      `json:"marks"`
	SideToPlay user.Side   `json:"sideToPlay"`
}

type SyncResult struct {
	ID         game.ID        `json:"id"`
	Users      []user.User    `json:"users"`
	Moves      []game.Move    `json:"moves"`
	State      string         `json:"state"`
	Marks      string         `json:"marks"`
	SideToPlay user.Side      `json:"sideToPlay,omitempty"`
	Winner     *winner.Winner `json:"winner,omitempty"`
}

type MoveResult struct {
	ID         game.ID   `json:"id"`
	Move       game.Move `json:"move"`
	State      string    `json:"state"`
	Marks      string    `json:"marks"`
	SideToPlay user.Side `json:"sideToPlay"`
}

type EndedResult struct {
	ID      game.ID        `json:"id"`
	Users   []user.User    `json:"users"`
	Moves   []game.Move    `json:"moves"`
	Marks   string         `json:"marks"`
	Outcome game.Outcome   `json:"outcome"`
	Winner  *winner.Winner `json:"winner,omitempty"`
}

type ErrorResult struct {
	UserID user.ID `json:"-"`
	Code   int     `json:"code"`
	Reason string  `json:"reason"`
}
