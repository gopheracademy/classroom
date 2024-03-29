// Copyright 2015 The Go Authors.  All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// +build appengine appenginevm

package main

import (
	"net/url"

	"github.com/gophertrain/classroom/present"

	_ "github.com/gophertrain/classroom/playground"
)

func initPlayground(basepath string, origin *url.URL) {
	playScript(basepath, "HTTPTransport")
}

func playable(c present.Code) bool {
	return present.PlayEnabled && c.Play && c.Ext == ".go"
}
