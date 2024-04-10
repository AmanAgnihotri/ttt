// Copyright 2024 Aman Agnihotri
// SPDX-License-Identifier: AGPL-3.0-only

package noun

import (
	"crypto/rand"
	"math/big"
	"strings"

	_ "embed"
)

//go:embed nouns.txt
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

type Noun string

func NewNoun() (Noun, bool) {
	if index, err := rand.Int(rand.Reader, count); err == nil {
		return Noun(Values[index.Int64()]), true
	}

	return "", false
}
