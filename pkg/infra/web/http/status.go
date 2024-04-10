// Copyright 2024 Aman Agnihotri
// SPDX-License-Identifier: AGPL-3.0-only

package http

type StatusCode int

const (
	StatusOk                   StatusCode = 200
	StatusNoContent            StatusCode = 204
	StatusNotModified          StatusCode = 304
	StatusBadRequest           StatusCode = 400
	StatusUnauthorized         StatusCode = 401
	StatusPaymentRequired      StatusCode = 402
	StatusForbidden            StatusCode = 403
	StatusNotFound             StatusCode = 404
	StatusMethodNotAllowed     StatusCode = 405
	StatusConflict             StatusCode = 409
	StatusGone                 StatusCode = 410
	StatusUnsupportedMediaType StatusCode = 415
	StatusLocked               StatusCode = 423
	StatusUpgradeRequired      StatusCode = 426
	StatusTooManyRequests      StatusCode = 429
	StatusInternalServerError  StatusCode = 500
	StatusNotImplemented       StatusCode = 501
	StatusBadGateway           StatusCode = 502
	StatusServiceUnavailable   StatusCode = 503
)
