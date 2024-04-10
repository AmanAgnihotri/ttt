// Copyright 2024 Aman Agnihotri
// SPDX-License-Identifier: AGPL-3.0-only

package jwt

import (
	"errors"

	"ttt/pkg/api/jwt/internal/audience"
	"ttt/pkg/api/jwt/internal/issuer"
	"ttt/pkg/api/jwt/internal/secret"
)

var (
	ErrInvalidAudience = errors.New("invalid audience")
	ErrInvalidIssuer   = errors.New("invalid issuer")
	ErrInvalidSecret   = errors.New("invalid secret key")
)

type Config struct {
	audience audience.Audience
	issuer   issuer.Issuer
	secret   secret.Secret
}

func NewConfig(key, iss, aud string) (Config, error) {
	var config Config

	jwtAudience, ok := audience.Parse(aud)
	if !ok {
		return config, ErrInvalidAudience
	}

	jwtIssuer, ok := issuer.Parse(iss)
	if !ok {
		return config, ErrInvalidIssuer
	}

	jwtSecret, ok := secret.Parse(key)
	if !ok {
		return config, ErrInvalidSecret
	}

	config = Config{jwtAudience, jwtIssuer, jwtSecret}

	return config, nil
}
