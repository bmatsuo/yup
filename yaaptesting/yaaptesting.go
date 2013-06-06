// Copyright 2013, Bryan Matsuo. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// yaaptesting.go [created: Thu,  6 Jun 2013]

/*
Package yaaptesting is a package for testing yaap assertions. It is
primarily used internally but can be used for testing custom assertion
functions.

The primary function provided by yaaptesting is Mock(). It is used to craete
safely contained pseudo tests.

	func Gt(t yaap.Test, x, y int, msg ..interface{}) {
		yaap.Assert(t, 1, x > y,
			fmt.Sprintf("%d is no greater than %d; %s", x, y, fmt.Sprint(msg...)))
	}

	func TestMyAssert(t *testing.T) {
		rec := Mock(func(t yaaptesting.Test) { MyAssertFunction(t, 1, 2) })
		yaap.Assert(t, rec.HadFatal(), "unexpected assertion pass")

		rec = Mock(func(t yaaptesting.Test) { MyAssertFunction(t, 2, 1) })
		yaap.Assert(t, !rec.HadFatal(), "unexpected assertion failure")
	}

Yaaptesting is fully tested. Testing is done using the "testing" package because
the yaap package's tests rely on yaaptesting.
*/
package yaaptesting

import (
	"fmt"
)

// identical to yaap.Test but included here to eliminate the dependency
type Test interface {
	Error(v ...interface{})
	Fatal(v ...interface{})
}

type mockFatal int

func Mock(fn func(Test)) *Recorder {
	rec := new(Recorder)
	mock(rec, fn)
	return rec
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

type Recorder struct {
	Log []*LogMessage
}

func (test *Recorder) Error(v ...interface{}) {
	test.Log = append(test.Log, LogError(v...))
}

func (test *Recorder) Fatal(v ...interface{}) {
	test.Log = append(test.Log, LogFatal(v...))
	panic(mockFatal(0))
}

func (test *Recorder) HadFatal() bool {
	if len(test.Log) == 0 {
		return false
	}
	return test.Log[len(test.Log)-1].Fatal()
}

type LogMessage struct {
	Typ   string // "error" or "fatal"
	Value string
}

func LogError(v ...interface{}) *LogMessage {
	return &LogMessage{"error", fmt.Sprint(v...)}
}

func LogFatal(v ...interface{}) *LogMessage {
	return &LogMessage{"fatal", fmt.Sprint(v...)}
}

func (msg *LogMessage) Fatal() bool {
	return msg.Typ == "fatal"
}
