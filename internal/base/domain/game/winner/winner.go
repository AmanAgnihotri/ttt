// Copyright 2024 Aman Agnihotri
// SPDX-License-Identifier: AGPL-3.0-only

package winner

type ID int64

type Winner struct {
	ID     ID     `bson:"id"     json:"id"`
	Method Method `bson:"method" json:"method"`
}
