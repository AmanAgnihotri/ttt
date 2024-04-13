// Copyright 2024 Aman Agnihotri
// SPDX-License-Identifier: AGPL-3.0-only

package api

import (
	"ttt/internal/base/domain/game/user"
	"ttt/internal/game/play/create/domain"
)

type Service interface {
	StartedResult() StartedResult
	SyncResult() SyncResult
	MoveResult() MoveResult
	EndedResult() EndedResult

	OutOfTurnPlayResult(userID user.ID) ErrorResult
	AlreadyOccupiedResult(userID user.ID) ErrorResult
	NotImplementedResult(userID user.ID) ErrorResult
}

func NewService(game *domain.Game) Service {
	return &service{game}
}

type service struct {
	game *domain.Game
}

func (s *service) StartedResult() StartedResult {
	return StartedResult{
		ID:         s.game.ID,
		Users:      s.game.Users,
		State:      s.game.StateAsString(),
		Marks:      s.game.MarksAsString(),
		SideToPlay: s.game.SideToPlay(),
	}
}

func (s *service) SyncResult() SyncResult {
	return SyncResult{
		ID:         s.game.ID,
		Users:      s.game.Users,
		Moves:      s.game.Moves,
		State:      s.game.StateAsString(),
		Marks:      s.game.MarksAsString(),
		SideToPlay: s.game.SideToPlay(),
		Winner:     s.game.Winner,
	}
}

func (s *service) MoveResult() MoveResult {
	return MoveResult{
		ID:         s.game.ID,
		Move:       s.game.LastMove(),
		State:      s.game.StateAsString(),
		Marks:      s.game.MarksAsString(),
		SideToPlay: s.game.SideToPlay(),
	}
}

func (s *service) EndedResult() EndedResult {
	return EndedResult{
		ID:      s.game.ID,
		Users:   s.game.Users,
		Moves:   s.game.Moves,
		Marks:   s.game.MarksAsString(),
		Outcome: s.game.Outcome,
		Winner:  s.game.Winner,
	}
}

func (s *service) OutOfTurnPlayResult(userID user.ID) ErrorResult {
	const forbidden = 403

	return ErrorResult{
		UserID: userID,
		Code:   forbidden,
		Reason: "out of turn play attempt",
	}
}

func (s *service) AlreadyOccupiedResult(userID user.ID) ErrorResult {
	const badRequest = 404

	return ErrorResult{
		UserID: userID,
		Code:   badRequest,
		Reason: "already occupied position",
	}
}

func (s *service) NotImplementedResult(userID user.ID) ErrorResult {
	const notImplemented = 501

	return ErrorResult{
		UserID: userID,
		Code:   notImplemented,
		Reason: "command not implemented",
	}
}
