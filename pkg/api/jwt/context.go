// Copyright 2024 Aman Agnihotri
// SPDX-License-Identifier: AGPL-3.0-only

package jwt

import (
	"errors"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"

	"ttt/pkg/api/jwt/internal/nonce"
)

const (
	Issuer    = "iss"
	Audience  = "aud"
	IssuedAt  = "iat"
	ExpiresAt = "exp"
	NotBefore = "nbf"
	ID        = "jti"
)

type Claims = map[string]any

type Context struct {
	secret   []byte
	issuer   string
	audience []string
	skew     time.Duration
}

var (
	ErrInvalidToken         = errors.New("invalid token")
	ErrInvalidSigningMethod = errors.New("invalid signing method")
)

func NewContext(c Config) *Context {
	return &Context{
		secret:   []byte(c.secret),
		issuer:   string(c.issuer),
		audience: []string{string(c.audience)},
		skew:     time.Minute,
	}
}

func (c *Context) NewClaims(time, expiryTime time.Time) (Claims, error) {
	randomID, err := nonce.NewNonce()
	if err != nil {
		return nil, fmt.Errorf("error creating claims: %w", err)
	}

	claims := jwt.MapClaims{
		Issuer:    c.issuer,
		Audience:  c.audience,
		IssuedAt:  jwt.NewNumericDate(time),
		ExpiresAt: jwt.NewNumericDate(expiryTime),
		NotBefore: jwt.NewNumericDate(time.Add(-c.skew)),
		ID:        string(randomID),
	}

	return claims, nil
}

func (c *Context) NewTokenFromClaims(claims Claims) (string, error) {
	var jwtClaims jwt.MapClaims = claims

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwtClaims)

	signedToken, err := token.SignedString(c.secret)
	if err != nil {
		return "", fmt.Errorf("error signing JWT: %w", err)
	}

	return signedToken, nil
}

func (c *Context) NewClaimsFromToken(t string) (Claims, error) {
	token, err := jwt.Parse(t, c.key)
	if err != nil {
		return nil, fmt.Errorf("error parsing token: %w", err)
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return claims, nil
	}

	return nil, ErrInvalidToken
}

func (c *Context) key(t *jwt.Token) (any, error) {
	if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
		return nil, fmt.Errorf("%w: %v", ErrInvalidSigningMethod, t.Header["alg"])
	}

	return c.secret, nil
}
