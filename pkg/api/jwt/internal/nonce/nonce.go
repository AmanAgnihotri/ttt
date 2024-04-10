// Copyright 2024 Aman Agnihotri
// SPDX-License-Identifier: AGPL-3.0-only

package nonce

import (
	"crypto/rand"
	"fmt"
	"math/big"
)

const (
	Length = 8
	chars  = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789ae"
)

var count = big.NewInt(int64(len(chars)))

type Nonce string

func NewNonce() (Nonce, error) {
	nonce := make([]byte, 0, Length)

	for range Length {
		if index, err := rand.Int(rand.Reader, count); err == nil {
			nonce = append(nonce, chars[index.Int64()])
		} else {
			return "", fmt.Errorf("error creating nonce: %w", err)
		}
	}

	return Nonce(nonce), nil
}
