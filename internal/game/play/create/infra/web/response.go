// Copyright 2024 Aman Agnihotri
// SPDX-License-Identifier: AGPL-3.0-only

package web

type Type string

const (
	Start Type = "start"
	Sync  Type = "sync"
	Move  Type = "move"
	End   Type = "end"
	Error Type = "error"
)

type Response[T any] struct {
	Type Type `json:"type"`
	Data T    `json:"data,omitempty"`
}
