// Copyright 2024 Aman Agnihotri
// SPDX-License-Identifier: AGPL-3.0-only

package lib

import (
	"fmt"

	"ttt/pkg/api/jwt"
)

const (
	secretKey   = "JWT_Secret"
	issuerKey   = "JWT_Issuer"
	audienceKey = "JWT_Audience"
)

func NewJWTContext() (*jwt.Context, error) {
	config, err := newJWTConfig()
	if err != nil {
		return nil, fmt.Errorf("error creating jwt config: %w", err)
	}

	jwtContext := jwt.NewContext(config)

	return jwtContext, nil
}

func newJWTConfig() (jwt.Config, error) {
	var config jwt.Config

	secret, err := LookupEnv(secretKey)
	if err != nil {
		return config, err
	}

	issuer, err := LookupEnv(issuerKey)
	if err != nil {
		return config, err
	}

	audience, err := LookupEnv(audienceKey)
	if err != nil {
		return config, err
	}

	config, err = jwt.NewConfig(secret, issuer, audience)
	if err != nil {
		return config, fmt.Errorf("error creating jwt config: %w", err)
	}

	return config, nil
}
