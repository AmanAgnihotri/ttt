// Copyright 2024 Aman Agnihotri
// SPDX-License-Identifier: AGPL-3.0-only

package web

type request struct {
	UserID int64  `json:"-"`
	Side   string `json:"side"`
}
