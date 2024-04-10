// Copyright 2024 Aman Agnihotri
// SPDX-License-Identifier: AGPL-3.0-only

package id

import (
	"crypto/rand"
	"encoding/binary"
	"math"
)

const (
	MinID int64 = 0x00038D7EA4C68000 // 10^15
	MaxID int64 = 0x001FFFFFFFFFFFFF // 2^53 - 1
)

func NewID() (int64, bool) {
	var bytes [8]byte
	if _, err := rand.Read(bytes[:]); err != nil {
		return 0, false
	}

	randomID := binary.BigEndian.Uint64(bytes[:])

	positiveID := int64(randomID & math.MaxInt64)

	clampedID := MinID + (positiveID % (MaxID - MinID + 1))

	return clampedID, true
}

func IsValid(i int64) bool {
	return i >= MinID && i < MaxID
}
