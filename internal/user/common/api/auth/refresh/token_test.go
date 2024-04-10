// Copyright 2024 Aman Agnihotri
// SPDX-License-Identifier: AGPL-3.0-only

package refresh_test

import (
	"testing"

	"ttt/internal/base/domain/user"
	"ttt/internal/base/domain/user/auth"
	"ttt/internal/user/common/api/auth/refresh"
)

func TestNewToken(t *testing.T) {
	t.Parallel()

	if userID, ok := user.NewID(); !ok {
		t.Errorf("error in user.NewID()")
	} else if authToken, ok := auth.NewToken(); !ok {
		t.Errorf("error in auth.NewToken()")
	} else {
		refreshToken := refresh.NewToken(userID, authToken)
		if len(string(refreshToken)) != 32 {
			t.Errorf("invalid refresh token")
		}
	}
}

func TestToken_Split(t *testing.T) {
	t.Parallel()

	if userID, ok := user.NewID(); !ok {
		t.Errorf("error in user.NewID()")
	} else if authToken, ok := auth.NewToken(); !ok {
		t.Errorf("error in auth.NewToken()")
	} else {
		refreshToken := refresh.NewToken(userID, authToken)
		if id, token, splitOk := refreshToken.Split(); !splitOk {
			t.Errorf("error in splitting refresh token")
		} else if id != userID || token != authToken {
			t.Errorf("invalid refresh token split")
		}
	}
}
