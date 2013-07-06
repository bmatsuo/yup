// Copyright 2013, Bryan Matsuo. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// yuptesting.go [created: Thu,  6 Jun 2013]

/*
Yaaptesting is a package for testing yup assertion functions. It is primarily
used for internal subpackages but can be used for testing custom assertion
functions.

Usage

The primary function provided by yuptesting is Mock(). It creates a mock
yup.Test interface to safely make assertions against. Mock() returns an
ErrorLog type against which assertions can be made about assertions in the
mock.

	// A custom assertion function.
	func Gt(t yup.Test, x, y int) {
		yup.Assert(t, 1, x > y, fmt.Sprintf("%d no greater than %d; %s", x, y))
	}

	// A test for the custom assertion function.
	func TestGt(t *testing.T) {
		rec := Mock(func(t yuptesting.Test) { Gt(t, 1, 2) })
		yup.Assert(t, 0, rec.HadFatal(), "assertion passed")

		rec = Mock(func(t yuptesting.Test) { Gt(t, 2, 1) })
		yup.Assert(t, 0, !rec.HadFatal(), "assertion failed")
	}

Notes

Yaaptesting is fully tested. Testing is done using the "testing" package because
the yup package's tests rely on yuptesting.
*/
package yuptesting

import (
	"fmt"
)

// Identical to yup.Test. Duplicated here to eliminate the dependency.
type Test interface {
	Error(v ...interface{})
	Fatal(v ...interface{})
}

type mockFatal int

// Create a mock Test implementation that inserts logs into an ErrorLog
func Mock(fn func(Test)) *ErrorLog {
	errlog := new(ErrorLog)
	mock(errlog.recorder(), fn)
	return errlog
}

func mock(t Test, fn func(Test)) {
	defer func() {
		if v := recover(); v != nil {
			if _, ok := v.(mockFatal); ok {
				return
			}
			panic(v)
		}
	}()
	fn(t)
}

// A mock test's error log.
type ErrorLog struct {
	Log []*LogMessage
}

// Returns true iff the mock logged any fatal errors.
func (test *ErrorLog) HadFatal() bool {
	if len(test.Log) == 0 {
		return false
	}
	return test.Log[len(test.Log)-1].Fatal()
}

func (test *ErrorLog) recorder() *recorder {
	return (*recorder)(test)
}

type recorder ErrorLog

func (test *recorder) Error(v ...interface{}) {
	test.Log = append(test.Log, logError(v...))
}

func (test *recorder) Fatal(v ...interface{}) {
	test.Log = append(test.Log, logFatal(v...))
	panic(mockFatal(0))
}

type LogMessage struct {
	Typ   string // "error" or "fatal"
	Value string
}

func logError(v ...interface{}) *LogMessage {
	return &LogMessage{"error", fmt.Sprint(v...)}
}

func logFatal(v ...interface{}) *LogMessage {
	return &LogMessage{"fatal", fmt.Sprint(v...)}
}

func (msg *LogMessage) Fatal() bool {
	return msg.Typ == "fatal"
}
