package web

import (
	"ttt/internal/base/domain/game/user"
)

type Session struct {
	UserID     user.ID
	Controller *Controller
}
