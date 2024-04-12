// Copyright 2024 Aman Agnihotri
// SPDX-License-Identifier: AGPL-3.0-only

package login_test

import (
	"testing"

	"ttt/internal/base/domain/user"
	"ttt/internal/base/domain/user/guest"
	"ttt/internal/user/common/api/user/login"
)

func TestNewID(t *testing.T) {
	t.Parallel()

	if userID, ok := user.NewID(); !ok {
		t.Errorf("error in user.NewID()")
	} else if guestToken, ok := guest.NewToken(); !ok {
		t.Errorf("error in guest.NewToken()")
	} else {
		loginID := login.NewID(userID, guestToken)
		if len(string(loginID)) != 32 {
			t.Errorf("invalid login id")
		}
	}
}

func TestID_Split(t *testing.T) {
	t.Parallel()

	if userID, ok := user.NewID(); !ok {
		t.Errorf("error in user.NewID()")
	} else if guestToken, ok := guest.NewToken(); !ok {
		t.Errorf("error in guest.NewToken()")
	} else {
		loginID := login.NewID(userID, guestToken)
		if id, token, splitOk := loginID.Split(); !splitOk {
			t.Errorf("error in splitting login id")
		} else if id != userID || token != guestToken {
			t.Errorf("invalid login id split")
		}
	}
}
