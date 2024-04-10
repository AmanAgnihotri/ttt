// Copyright 2024 Aman Agnihotri
// SPDX-License-Identifier: AGPL-3.0-only

package user

import (
	"crypto/rand"
)

const (
	random = "random"
	xSide  = "x"
	oSide  = "o"
)

type Side string

const (
	XSide Side = xSide
	OSide Side = oSide
)

func ParseSide(s string) (Side, bool) {
	switch s {
	case random, xSide, oSide:
		return Side(s), true
	default:
		return "", false
	}
}

func (s Side) IsValid() bool {
	return s == xSide || s == oSide
}

func (s Side) Resolve() Side {
	switch s {
	case xSide, oSide:
		return s
	default:
		return s.random()
	}
}

func (s Side) Opposite() Side {
	switch s {
	case xSide:
		return oSide
	case oSide:
		return xSide
	default:
		panic("invalid side")
	}
}

func (s Side) random() Side {
	bytes := make([]byte, 1)

	_, err := rand.Read(bytes)
	if err != nil {
		return xSide
	}

	bit := bytes[0] & 1

	if bit == 0 {
		return xSide
	}

	return oSide
}
