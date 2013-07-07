// Copyright 2013, Bryan Matsuo. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// yup.go [created: Wed,  5 Jun 2013]

/*
Yet another assertion package. Simple and extensible.

The yup package exports a single generic function, Assert(). More complex
assertions are defined in the package's subdirectories.
*/
package yup

import (
	"fmt"
	"path/filepath"
	"runtime"
)

// when true, assertion failure locations are logged on the line preceeding
// the failure message. when false, carriage returns are used to overwrite the
// location information included by the testing package.
var CompatabilityMode = false

// returns the nth caller in getcaller()'s grandcaller's call stack. when n
// is 0 getcaller() returns its grandcaller's location.
func getcaller(n int) (file string, line int) {
	_, file, line, ok := runtime.Caller(2 + n)
	if ok {
		return file, line
	}
	return "", -1
}

func T(t Test, n int, ok bool, msg ...interface{}) {
	if !ok {
		file, line := getcaller(n)
		caller := fmt.Sprintf("%s:%d", filepath.Base(file), line)
		if CompatabilityMode {
			// don't do anything crazy
			t.Error(caller)
			t.Fatal(fmt.Sprint(msg...))
			return // in case t is a weird implementation
		}
		// testing package hack. override line number
		t.Fatal(fmt.Sprintf("\r\t%s: ", caller), fmt.Sprint(msg...))
	}
}

// a generic assertion function. if ok is false then msg is logged as a fatal
// error. the line number logged with msg is depth-th caller in Assert()'s
// caller's call stack. a depth of zero logs the location of Assert()'s caller.
func Assert(t Test, ok bool, msg ...interface{}) {
	T(t, 1, ok, msg...)
}

// a minimal interface implemented by Test
type Test interface {
	Error(v ...interface{})
	Fatal(v ...interface{})
}
