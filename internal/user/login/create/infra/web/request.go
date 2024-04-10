// Copyright 2024 Aman Agnihotri
// SPDX-License-Identifier: AGPL-3.0-only

package web

type request struct {
	Platform   string `json:"platform"`
	AppVersion int32  `json:"appVersion"`
	LoginID    string `json:"loginId"`
}
