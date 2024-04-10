// Copyright 2024 Aman Agnihotri
// SPDX-License-Identifier: AGPL-3.0-only

package adjective

import (
	"crypto/rand"
	"math/big"
	"strings"

	_ "embed"
)

//go:embed adjectives.txt
var text string

var (
	Values []string
	count  *big.Int
)

func init() {
	Values = strings.Split(text, "\n")
	count = big.NewInt(int64(len(Values)))

	text = ""
}

type Adjective string

func NewAdjective() (Adjective, bool) {
	if index, err := rand.Int(rand.Reader, count); err == nil {
		return Adjective(Values[index.Int64()]), true
	}

	return "", false
}
