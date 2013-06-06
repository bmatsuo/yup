// Copyright 2013, Bryan Matsuo. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// yaap.go [created: Wed,  5 Jun 2013]

// Package yaap does ....
package yaap

import (
	"fmt"
	"path/filepath"
	"runtime"
)

var CompatabilityMode = false

// returns the nth caller in getcaller()'s caller's
// call stack
func getcaller(n int) (file string, line int) {
	_, file, line, ok := runtime.Caller(2+n)
	if ok {
		return file, line
	}
	return "", -1
}

func T(t Test, ok bool, msg ...interface{}) {
	Assert(t, 1, ok, msg...)
}

// this should probably be public, but i need a better name
func Assert(t Test, depth int, ok bool, msg ...interface{}) {
	if !ok {
		file, line := getcaller(depth)
		caller := fmt.Sprintf("%s:%d", filepath.Base(file), line)
		if CompatabilityMode {
			// don't do anything crazy
			t.Error(caller)
			t.Fatal(fmt.Sprint(msg...))
		}
		// testing package hack
		t.Fatal(fmt.Sprint("\r\t%s: ", caller), fmt.Sprint(msg...))
	}
}

// a minimal interface implemented by Test
type Test interface {
	Error(v ...interface{})
	Fatal(v ...interface{})
}
