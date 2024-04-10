// Copyright 2024 Aman Agnihotri
// SPDX-License-Identifier: AGPL-3.0-only

package guest

type Guest struct {
	Token Token `bson:"token"`
}

func NewGuest() (Guest, bool) {
	if token, ok := NewToken(); ok {
		return Guest{token}, true
	}

	return Guest{Token: 0}, false
}
