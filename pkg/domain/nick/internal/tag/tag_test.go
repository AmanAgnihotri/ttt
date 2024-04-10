// Copyright 2024 Aman Agnihotri
// SPDX-License-Identifier: AGPL-3.0-only

package tag_test

import (
	"strconv"
	"testing"

	"ttt/pkg/domain/nick/internal/tag"
)

func TestNewTag(t *testing.T) {
	t.Parallel()

	if newTag, ok := tag.NewTag(); !ok {
		t.Errorf("tag.NewTag() threw error")
	} else if !isValid(newTag, tag.MaxTagExclusive) {
		t.Errorf("invalid tag: %s", newTag)
	}
}

func isValid(tag tag.Tag, maxTagExclusive int) bool {
	tagAsString := string(tag)

	maxTag := strconv.Itoa(maxTagExclusive - 1)
	if len(tagAsString) != len(maxTag) {
		return false
	}

	tagAsInt, err := strconv.Atoi(tagAsString)
	if err != nil {
		return false
	}

	return tagAsInt >= 0 && tagAsInt < maxTagExclusive
}
