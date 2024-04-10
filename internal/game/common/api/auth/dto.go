// Copyright 2024 Aman Agnihotri
// SPDX-License-Identifier: AGPL-3.0-only

package auth

type DTO struct {
	TokenType        string `json:"tokenType"`
	AccessToken      string `json:"accessToken"`
	ExpiresInSeconds int    `json:"expiresInSeconds"`
}
