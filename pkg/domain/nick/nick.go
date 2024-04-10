// Copyright 2024 Aman Agnihotri
// SPDX-License-Identifier: AGPL-3.0-only

package nick

import (
	"fmt"

	"ttt/pkg/domain/nick/internal/adjective"
	"ttt/pkg/domain/nick/internal/noun"
	"ttt/pkg/domain/nick/internal/tag"
)

func NewNick() (string, bool) {
	newAdjective, ok := adjective.NewAdjective()
	if !ok {
		return "", false
	}

	newNoun, ok := noun.NewNoun()
	if !ok {
		return "", false
	}

	newTag, ok := tag.NewTag()
	if !ok {
		return "", false
	}

	newNick := fmt.Sprintf("%s%s%s", newAdjective, newNoun, newTag)

	return newNick, true
}
