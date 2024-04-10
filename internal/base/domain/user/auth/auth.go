// Copyright 2024 Aman Agnihotri
// SPDX-License-Identifier: AGPL-3.0-only

package auth

import "time"

type Auth struct {
	Token      Token     `bson:"token"`
	ExpiryTime time.Time `bson:"expiryTime"`
}

func NewAuth(createTime time.Time) (*Auth, bool) {
	if token, ok := NewToken(); ok {
		expiryTime := createTime.Add(7 * 24 * time.Hour)

		return &Auth{token, expiryTime}, true
	}

	return nil, false
}

func (a *Auth) IsValid(time time.Time) bool {
	return a != nil && a.ExpiryTime.After(time)
}
