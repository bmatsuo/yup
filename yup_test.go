// Copyright 2013, Bryan Matsuo. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// yup_test.go [created: Thu,  6 Jun 2013]

package yup

import (
	"github.com/bmatsuo/yup/yuptesting"

	"fmt"
	"runtime"
	"strings"
	"testing"
)

func T(t Test, ok bool, msg ...interface{}) {
	Assert(t, 1, ok, msg...)
}

func TestAssert(t *testing.T) {
	var preAssertLine int
	rec := yuptesting.Mock(func(t yuptesting.Test) {
		Assert(t, 0, true, "a passed assertion")

		_, _, preAssertLine, _ = runtime.Caller(0) // this is a little crazy
		Assert(t, 0, false, "a failed assertion")

		Assert(t, 0, false, "this failure is never seen")
	})
	if len(rec.Log) != 1 {
		t.Errorf("expected 1 log message but fonud %d", len(rec.Log))
	}
	last := rec.Log[len(rec.Log)-1]
	if !last.Fatal() {
		t.Fatal("unexpected non-fatal error")
	}
	expectedPos := fmt.Sprintf("%s:%d", "yup_test.go", preAssertLine+1)
	if -1 == strings.Index(last.Value, expectedPos) {
		t.Fatalf("position %q not found %q", expectedPos, last.Value)
	}
}

func TestCompatabilityMode(t *testing.T) {
	CompatabilityMode = true
	defer func() { CompatabilityMode = false }()
	var preAssertLine int
	rec := yuptesting.Mock(func(t yuptesting.Test) {
		_, _, preAssertLine, _ = runtime.Caller(0) // this is a little crazy
		Assert(t, 0, false, "a failed assertion")
	})
	if len(rec.Log) != 2 {
		t.Errorf("expected 2 log messages but found %d", len(rec.Log))
	}
	last := rec.Log[len(rec.Log)-2]
	if last.Fatal() {
		t.Fatal("unexpected fatal error")
	}
	expectedPos := fmt.Sprintf("%s:%d", "yup_test.go", preAssertLine+1)
	if -1 == strings.Index(last.Value, expectedPos) {
		t.Fatalf("position %q not found %q", expectedPos, last.Value)
	}
}

func TestT(t *testing.T) {
	rec := yuptesting.Mock(func(t yuptesting.Test) { T(t, false) })
	Assert(t, 0, rec.HadFatal(), "unexpected assertion pass")

	rec = yuptesting.Mock(func(t yuptesting.Test) { T(t, true) })
	Assert(t, 0, !rec.HadFatal(), "unexpected assertion failure")
}