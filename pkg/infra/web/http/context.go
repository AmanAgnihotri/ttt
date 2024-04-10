// Copyright 2024 Aman Agnihotri
// SPDX-License-Identifier: AGPL-3.0-only

package http

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
)

type key byte

const (
	claimsKey key = iota
)

type Context struct {
	context.Context
	r *http.Request
	http.ResponseWriter
}

func NewContext(
	c context.Context,
	r *http.Request,
	w http.ResponseWriter,
) *Context {
	return &Context{c, r, w}
}

func (c *Context) Claims() (map[string]any, bool) {
	claims, ok := c.Value(claimsKey).(map[string]any)

	return claims, ok
}

func (c *Context) ClaimValue(name string) (any, bool) {
	claims, ok := c.Claims()
	if !ok {
		return "", false
	}

	value, ok := claims[name]

	return value, ok
}

func (c *Context) ClaimValueAsString(name string) (string, bool) {
	value, ok := c.ClaimValue(name)
	if !ok {
		return "", false
	}

	valueAsString, ok := value.(string)

	return valueAsString, ok
}

func (c *Context) ClaimValueAsFloat64(name string) (float64, bool) {
	value, ok := c.ClaimValue(name)
	if !ok {
		return 0, false
	}

	valueAsFloat64, ok := value.(float64)

	return valueAsFloat64, ok
}

func (c *Context) ClaimValueAsInt64(name string) (int64, bool) {
	if value, ok := c.ClaimValueAsFloat64(name); ok {
		return int64(value), true
	}

	return 0, false
}

func (c *Context) PathValue(name string) string {
	return c.r.PathValue(name)
}

func (c *Context) HeaderValue(key string) string {
	return c.r.Header.Get(key)
}

func (c *Context) Body(v any) error {
	err := json.NewDecoder(c.r.Body).Decode(v)
	if err != nil {
		return fmt.Errorf("error reading body: %w", err)
	}

	return nil
}

func (c *Context) Ok(data any) {
	c.Data(StatusOk, data)
}

func (c *Context) PlainText(text string) {
	c.Header().Set("Content-Type", "text/plain; charset=utf-8")
	c.WriteHeader(int(StatusOk))
	_, _ = c.Write([]byte(text))
}

func (c *Context) Unauthorized() {
	c.Status(StatusUnauthorized)
}

func (c *Context) BadRequest() {
	c.Status(StatusBadRequest)
}

func (c *Context) Data(code StatusCode, data any) {
	c.Header().Set("Content-Type", "application/json; charset=utf-8")
	c.WriteHeader(int(code))

	err := json.NewEncoder(c).Encode(data)
	if err != nil {
		return
	}
}

func (c *Context) Status(code StatusCode) {
	c.WriteHeader(int(code))
}

func (c *Context) Cached(data any, knownEtag string) {
	etag := c.r.Header.Get("If-None-Match")

	if etag == knownEtag {
		c.WriteHeader(int(StatusNotModified))
	} else {
		c.Header().Add("ETag", knownEtag)
		c.Header().Add("Cache-Control", "public, no-cache")

		err := json.NewEncoder(c).Encode(data)
		if err != nil {
			return
		}
	}
}

func (c *Context) req() *http.Request {
	return c.r
}

func (c *Context) rw() http.ResponseWriter {
	return c.ResponseWriter
}
