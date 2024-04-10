// Copyright 2024 Aman Agnihotri
// SPDX-License-Identifier: AGPL-3.0-only

package tag

import (
	"crypto/rand"
	"fmt"
	"math/big"
)

const (
	MaxTagExclusive = 10_000
)

var count = big.NewInt(MaxTagExclusive)

type Tag string

func NewTag() (Tag, bool) {
	if index, err := rand.Int(rand.Reader, count); err == nil {
		return Tag(fmt.Sprintf("%04d", index.Int64())), true
	}

	return "", false
}
