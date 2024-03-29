// Copyright 2012 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// +build !appengine

package main

// This package registers "/compile" and "/share" handlers
// that redirect to the golang.org playground.
import _ "github.com/gophertrain/classroom/playground"
