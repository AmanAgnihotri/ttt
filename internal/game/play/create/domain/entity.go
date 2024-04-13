// Copyright 2024 Aman Agnihotri
// SPDX-License-Identifier: AGPL-3.0-only

package domain

import (
	"fmt"
	"time"

	"ttt/internal/base/domain/game"
	"ttt/internal/base/domain/game/user"
	"ttt/internal/base/domain/game/winner"
	"ttt/internal/game/common/domain/board"
)

type Game struct {
	ID         game.ID        `bson:"_id"`
	FinishTime time.Time      `bson:"finishTime"`
	Outcome    game.Outcome   `bson:"outcome"`
	Winner     *winner.Winner `bson:"winner,omitempty"`
	Users      []user.User    `bson:"users"`
	Moves      []game.Move    `bson:"moves"`

	board   *board.Board
	current *user.User
}

func (g *Game) IsPlayable() bool {
	const maxUserCount = 2

	return g != nil && len(g.Users) == maxUserCount
}

func (g *Game) Init() {
	g.board = board.NewBoard()

	for i, u := range g.Users {
		if u.Side == user.XSide {
			g.current = &g.Users[i]

			break
		}
	}
}

func (g *Game) Apply(play Play) Event {
	if g.current == nil {
		return Unplayable
	}

	if g.current.ID != play.UserID {
		return OutOfTurnPlay
	}

	position := play.Position
	mark := g.getCurrentMark()

	if ok := g.board.Apply(position, mark); !ok {
		return AlreadyOccupied
	}

	g.addToMoves(position, mark)

	switch g.board.State() {
	case board.Playable:
		g.current = g.getOtherUser()

		return Applied

	case board.Draw:
		g.Outcome = game.Draw
		g.FinishTime = play.Time
		g.current = nil

		return Completed

	default:
		g.updateAsWin(play.Time, winner.Play)

		return Completed
	}
}

func (g *Game) TimedOut(time time.Time) {
	g.updateAsWin(time, winner.Timeout)
}

func (g *Game) LastMove() game.Move {
	if len(g.Moves) > 0 {
		return g.Moves[len(g.Moves)-1]
	}

	return ""
}

func (g *Game) StateAsString() string {
	return g.board.State().String()
}

func (g *Game) MarksAsString() string {
	return string(g.board.Marks())
}

func (g *Game) SideToPlay() user.Side {
	if g.current != nil {
		return g.current.Side
	}

	return ""
}

func (g *Game) updateAsWin(time time.Time, method winner.Method) {
	if g.current == nil {
		return
	}

	g.Outcome = game.Win
	g.FinishTime = time

	var winnerID winner.ID
	if method == winner.Play {
		winnerID = winner.ID(g.current.ID)
	} else {
		winnerID = winner.ID(g.getOtherUser().ID)
	}

	g.Winner = &winner.Winner{
		ID:     winnerID,
		Method: method,
	}

	g.current = nil
}

func (g *Game) getCurrentMark() board.Mark {
	if g.current.Side == user.XSide {
		return board.XMark
	}

	return board.OMark
}

func (g *Game) addToMoves(position board.Position, mark board.Mark) {
	move := game.Move(fmt.Sprintf("%c%d", mark, position))

	g.Moves = append(g.Moves, move)
}

func (g *Game) getOtherUser() *user.User {
	if g.current == nil {
		return nil
	}

	for _, u := range g.Users {
		if u.ID != g.current.ID {
			return &u
		}
	}

	return nil
}
