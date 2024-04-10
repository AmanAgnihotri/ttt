// Copyright 2024 Aman Agnihotri
// SPDX-License-Identifier: AGPL-3.0-only

package user

type State string

const (
	Deleted State = "deleted"
	Banned  State = "banned"
	Locked  State = "locked"
	Guest   State = "guest"
)
