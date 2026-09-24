// Copyright 2026 Aman Agnihotri
// SPDX-License-Identifier: AGPL-3.0-only

package domain_test

import (
	"testing"
	"time"

	"ttt/internal/base/domain/game"
	"ttt/internal/base/domain/game/user"
	"ttt/internal/base/domain/game/winner"
	"ttt/internal/game/common/domain/board"
	"ttt/internal/game/play/create/domain"
)

const (
	xID = user.ID(1)
	oID = user.ID(2)
)

func TestTurnOrder(t *testing.T) {
	t.Parallel()

	newGame := createGame()

	if event := newGame.Apply(newPlay(oID, 4)); event != domain.OutOfTurnPlay {
		t.Errorf("event = %v; expected %v", event, domain.OutOfTurnPlay)
	}

	if event := newGame.Apply(newPlay(xID, 4)); event != domain.Applied {
		t.Errorf("event = %v; expected %v", event, domain.Applied)
	}

	if event := newGame.Apply(newPlay(oID, 4)); event != domain.AlreadyOccupied {
		t.Errorf("event = %v; expected %v", event, domain.AlreadyOccupied)
	}
}

func TestWin(t *testing.T) {
	t.Parallel()

	newGame := createGame()
	userIDs := []user.ID{xID, oID}
	event := domain.Unplayable

	for i, position := range []board.Position{0, 3, 1, 4, 2} {
		event = newGame.Apply(newPlay(userIDs[i%2], position))
	}

	if event != domain.Completed || newGame.Outcome != game.Win {
		t.Fatalf(
			"event, outcome = %v, %v; expected %v, %v",
			event, newGame.Outcome, domain.Completed, game.Win)
	}

	if !isWinner(newGame, xID, winner.Play) {
		t.Errorf("winner = %+v; expected %v by %v",
			newGame.Winner, xID, winner.Play)
	}

	if event = newGame.Apply(newPlay(oID, 5)); event != domain.Unplayable {
		t.Errorf("event = %v; expected %v", event, domain.Unplayable)
	}
}

func TestDraw(t *testing.T) {
	t.Parallel()

	newGame := createGame()
	userIDs := []user.ID{xID, oID}

	for i, position := range []board.Position{0, 1, 2, 4, 3, 5, 7, 6, 8} {
		newGame.Apply(newPlay(userIDs[i%2], position))
	}

	if newGame.Outcome != game.Draw || newGame.Winner != nil {
		t.Errorf(
			"outcome, winner = %v, %+v; expected %v, nil",
			newGame.Outcome, newGame.Winner, game.Draw)
	}
}

func TestTimedOut(t *testing.T) {
	t.Parallel()

	newGame := createGame()

	newGame.Apply(newPlay(xID, 4))
	newGame.TimedOut(time.Now())

	if !isWinner(newGame, xID, winner.Timeout) {
		t.Errorf("winner = %+v; expected %v by %v",
			newGame.Winner, xID, winner.Timeout)
	}
}

func createGame() *domain.Game {
	newGame := &domain.Game{
		ID:         game.ID(1),
		FinishTime: time.Time{},
		Outcome:    "",
		Winner:     nil,
		Users: []user.User{
			{ID: xID, Nick: "x", Side: user.XSide},
			{ID: oID, Nick: "o", Side: user.OSide},
		},
		Moves: nil,
	}

	newGame.Init()

	return newGame
}

func newPlay(userID user.ID, position board.Position) domain.Play {
	return domain.Play{UserID: userID, Position: position, Time: time.Now()}
}

func isWinner(g *domain.Game, userID user.ID, method winner.Method) bool {
	w := g.Winner

	return w != nil && w.ID == winner.ID(userID) && w.Method == method
}
