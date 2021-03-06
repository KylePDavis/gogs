// Copyright 2014 The Gogs Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package base

type (
	TplName string

	ApiJsonErr struct {
		Message string `json:"message"`
		DocUrl  string `json:"url"`
	}
)

var GoGetMetas = make(map[string]bool)
