// Copyright 2013, Bryan Matsuo. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// yup.go [created: Wed,  5 Jun 2013]

/*
Yup, another assertion package. Simple and extensible.

The yup package exports simple, generic assertion functions that can be used
directly in tests. There are also low-level functions for defining high-level
assertions that reduce boilerplace such as those seen in the yup* subpackages.
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

// low-level function that issues a fatal error to t, reporting a position in
// the call stack to the user. if n is 0, the position reported to the user is
// the invocation of F that triggered the error.
func F(t Test, n uint, msg ...interface{}) {
	file, line := getcaller(int(n))
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

// low-level test function used to write other test functions.
// pass the call depth n which you want to be reported to the user
// if the test fails. see F().
func T(t Test, n uint, ok bool, msg ...interface{}) {
	if !ok {
		F(t, n+1, msg...)
	}
}

// test with a deferred computation. fn is called only if ok is false.
// see T().
func TD(t Test, n uint, ok bool, fn func() string) {
	if !ok {
		F(t, n+1, fn())
	}
}

// a simple assertion. see T().
func Assert(t Test, ok bool, msg ...interface{}) {
	T(t, 1, ok, msg...)
}

// an assertion with a deferred computation. see TD().
func AssertD(t Test, n uint, ok bool, fn func() string) {
	if !ok {
		F(t, n+1, fn())
	}
}

// fail the test. see F().
func Fail(t Test, msg ...interface{}) {
	F(t, 1, msg...)
}

// a minimal interface implemented by Test
type Test interface {
	Error(v ...interface{})
	Fatal(v ...interface{})
}
