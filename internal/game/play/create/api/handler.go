// Copyright 2024 Aman Agnihotri
// SPDX-License-Identifier: AGPL-3.0-only

package api

import (
	"time"

	"ttt/internal/game/play/create/domain"
)

const (
	bufferSize         = 2
	maxDurationPerMove = 60 * time.Second
)

type Handler interface {
	Handle(command domain.Command)
	Loop() <-chan Result
	Sync() SyncResult
	Game() *domain.Game
}

func NewHandler(game *domain.Game) Handler {
	game.Init()

	return &handler{
		incoming: make(chan domain.Command, bufferSize),
		game:     game,
		Service:  NewService(game),
	}
}

type handler struct {
	incoming chan domain.Command
	game     *domain.Game
	Service
}

func (h *handler) Handle(command domain.Command) {
	h.incoming <- command
}

func (h *handler) Loop() <-chan Result {
	outgoing := make(chan Result, bufferSize)

	go h.loop(outgoing)

	return outgoing
}

func (h *handler) loop(outgoing chan<- Result) {
	outgoing <- h.StartedResult()

	ticker := time.NewTicker(maxDurationPerMove)

	defer func() {
		close(outgoing)
		ticker.Stop()
	}()

	for {
		select {
		case command := <-h.incoming:
			if play, ok := command.(domain.Play); ok {
				event := h.game.Apply(play)

				switch event {
				case domain.Applied:
					ticker.Reset(maxDurationPerMove)
					outgoing <- h.MoveResult()

				case domain.Completed, domain.Unplayable:
					ticker.Stop()
					outgoing <- h.EndedResult()

					return

				case domain.OutOfTurnPlay:
					outgoing <- h.OutOfTurnPlayResult(play.UserID)

				case domain.AlreadyOccupied:
					outgoing <- h.AlreadyOccupiedResult(play.UserID)
				}
			}

		case <-ticker.C:
			ticker.Stop()
			h.game.TimedOut(time.Now().UTC())
			outgoing <- h.EndedResult()

			return
		}
	}
}

func (h *handler) Sync() SyncResult {
	return h.SyncResult()
}

func (h *handler) Game() *domain.Game {
	return h.game
}
