// Copyright 2024 Aman Agnihotri
// SPDX-License-Identifier: AGPL-3.0-only

package board

import (
	"ttt/internal/base/domain/game/user"
)

type Mark byte

const (
	Empty Mark = '.'
	XMark Mark = 'x'
	OMark Mark = 'o'
)

func (m Mark) IsEmpty() bool {
	return m == Empty
}

func (m Mark) IsValid() bool {
	return m == XMark || m == OMark
}

func (m Mark) Opposite() Mark {
	switch m {
	case OMark:
		return XMark
	case XMark:
		return OMark
	default:
		return Empty
	}
}

func (m Mark) ToSide() user.Side {
	switch m {
	case XMark:
		return user.XSide
	case OMark:
		return user.OSide
	default:
		panic("invalid mark")
	}
}

func (m Mark) String() string {
	switch m {
	case OMark:
		return "o"
	case XMark:
		return "x"
	default:
		return ""
	}
}
