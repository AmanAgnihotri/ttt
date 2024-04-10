// Copyright 2024 Aman Agnihotri
// SPDX-License-Identifier: AGPL-3.0-only

package lib

import (
	"errors"
	"fmt"
	"os"
)

var ErrVariableNotSet = errors.New("environment variable not set")

func LookupEnv(key string) (string, error) {
	if value, ok := os.LookupEnv(key); ok {
		return value, nil
	}

	return "", fmt.Errorf("%w: %s", ErrVariableNotSet, key)
}
