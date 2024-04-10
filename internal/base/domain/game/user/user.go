// Copyright 2024 Aman Agnihotri
// SPDX-License-Identifier: AGPL-3.0-only

package user

type Nick string

type User struct {
	ID   ID   `bson:"id"   json:"id"`
	Nick Nick `bson:"nick" json:"nick"`
	Side Side `bson:"side" json:"side"`
}
