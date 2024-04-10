// Copyright 2024 Aman Agnihotri
// SPDX-License-Identifier: AGPL-3.0-only

package time

import "time"

type Timer interface {
	UTCNow() time.Time
}

func NewTimer() Timer {
	return &timer{}
}

type timer struct{}

func (t *timer) UTCNow() time.Time {
	return time.Now().UTC()
}
