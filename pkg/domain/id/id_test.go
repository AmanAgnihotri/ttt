// Copyright 2024 Aman Agnihotri
// SPDX-License-Identifier: AGPL-3.0-only

package id_test

import (
	"strconv"
	"testing"

	"ttt/pkg/domain/id"
)

func TestNewID(t *testing.T) {
	t.Parallel()

	if newID, ok := id.NewID(); !ok {
		t.Errorf("error in creating id")
	} else if !id.IsValid(newID) {
		t.Errorf("invalid id")
	} else {
		idAsString := strconv.FormatInt(newID, 10)
		if len(idAsString) != 16 {
			t.Errorf("invalid id length")
		}
	}
}
