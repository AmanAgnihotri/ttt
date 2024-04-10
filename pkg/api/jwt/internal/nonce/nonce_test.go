// Copyright 2024 Aman Agnihotri
// SPDX-License-Identifier: AGPL-3.0-only

package nonce_test

import (
	"testing"

	"ttt/pkg/api/jwt/internal/nonce"
)

func TestNewNonce(t *testing.T) {
	t.Parallel()

	result, err := nonce.NewNonce()
	if err != nil {
		t.Errorf("%v", err)
	} else if len(result) != nonce.Length {
		t.Errorf(
			"nonce.NewNonce() = %v, length: %d; expected length %d",
			result, len(result), nonce.Length)
	}
}
