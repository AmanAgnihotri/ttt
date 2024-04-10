// Copyright 2024 Aman Agnihotri
// SPDX-License-Identifier: AGPL-3.0-only

package game

import (
	"crypto/rand"
	"math/big"
	"strings"
)

const (
	length = 5
	chars  = "0123456789ABCDEFGHJKMNPQRSTVWXYZ"
)

var (
	count   = big.NewInt(int64(len(chars)))
	mapping = map[rune]byte{}
)

func init() {
	const extra = "ILO"

	mapping = make(map[rune]byte, len(chars)+len(extra))

	for _, c := range chars {
		mapping[c] = byte(c)
	}

	mapping['I'] = '1'
	mapping['L'] = '1'
	mapping['O'] = '0'
}

type Code string

func NewCode() (Code, bool) {
	code := make([]byte, 0, length)

	for range length {
		if index, err := rand.Int(rand.Reader, count); err == nil {
			code = append(code, chars[index.Int64()])
		} else {
			return "", false
		}
	}

	return Code(code), true
}

func ParseCode(code string) (Code, bool) {
	code = strings.ToUpper(code)

	builder := strings.Builder{}
	builder.Grow(len(code))

	for _, r := range code {
		if value, exists := mapping[r]; exists {
			builder.WriteByte(value)
		} else {
			return "", false
		}
	}

	return Code(builder.String()), true
}
