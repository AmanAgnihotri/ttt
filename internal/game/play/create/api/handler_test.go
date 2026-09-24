// Copyright 2026 Aman Agnihotri
// SPDX-License-Identifier: AGPL-3.0-only

package api_test

import (
	"sync"
	"testing"
	"time"

	"ttt/internal/base/domain/game"
	"ttt/internal/base/domain/game/user"
	"ttt/internal/base/domain/game/winner"
	"ttt/internal/game/common/domain/board"
	"ttt/internal/game/play/create/api"
	"ttt/internal/game/play/create/domain"
)

const (
	xID = user.ID(1)
	oID = user.ID(2)
)

func TestSyncWhilePlaying(t *testing.T) {
	t.Parallel()

	const syncCount = 100

	handler := api.NewHandler(createGame())
	last := lastResult(handler.Loop())

	var waitGroup sync.WaitGroup

	waitGroup.Go(func() {
		for range syncCount {
			handler.Sync()
		}
	})

	userIDs := []user.ID{xID, oID}

	for i, position := range []board.Position{0, 3, 1, 4, 2} {
		command := domain.Play{
			UserID:   userIDs[i%2],
			Position: position,
			Time:     time.Now(),
		}

		handler.Handle(command)
	}

	waitGroup.Wait()

	result := <-last

	ended, ok := result.(api.EndedResult)
	if !ok {
		t.Fatalf("result = %T; expected api.EndedResult", result)
	}

	if ended.Winner == nil || ended.Winner.ID != winner.ID(xID) {
		t.Errorf("winner = %+v; expected %v", ended.Winner, xID)
	}
}

func createGame() *domain.Game {
	return &domain.Game{
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
}

func lastResult(results <-chan api.Result) <-chan api.Result {
	last := make(chan api.Result, 1)

	go func() {
		var final api.Result

		for result := range results {
			final = result
		}

		last <- final
	}()

	return last
}
