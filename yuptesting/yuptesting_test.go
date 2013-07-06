// Copyright 2013, Bryan Matsuo. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// yuptesting_test.go [created: Thu,  6 Jun 2013]

package yuptesting

import (
	"testing"
)

func TestMockPass(t *testing.T) {
	rec := Mock(func(mocktest Test) {})
	if rec.HadFatal() {
		t.Fatalf("unexpected fatal error")
	}
	if len(rec.Log) != 0 {
		t.Fatalf("exected empty log but found %d messages", len(rec.Log))
	}
}

func TestMockError(t *testing.T) {
	rec := Mock(func(mocktest Test) {
		mocktest.Error("this is an error")
	})
	if rec.HadFatal() {
		t.Fatalf("unexpected fatal error")
	}
	if len(rec.Log) != 1 {
		t.Fatalf("exected 1 message but found %d", len(rec.Log))
	}
}

func TestMockFail(t *testing.T) {
	rec := Mock(func(mocktest Test) {
	})
	if len(rec.Log) != 0 {
		t.Fatalf("exected empty log but found %d messages", len(rec.Log))
	}
	rec = Mock(func(mocktest Test) {
		mocktest.Error("this is ", "an error")
		mocktest.Fatal("this is ", "fatal")
		mocktest.Fatal("this is ", "never seen")
	})
	if !rec.HadFatal() {
		t.Fatalf("expected a fatal error but did not have one")
	}
	if len(rec.Log) != 2 {
		t.Fatalf("expected 2 log messages but got %d", len(rec.Log))
	}
	if rec.Log[0].Fatal() {
		t.Fatalf("unexpected fatal error: %v", rec.Log[0])
	}
	if rec.Log[0].Value != "this is an error" {
		t.Fatalf("expected \"this is an error\" but got %q", rec.Log[0].Value)
	}
	if !rec.Log[1].Fatal() {
		t.Fatalf("unexpected non-fatal error", rec.Log[1])
	}
	if rec.Log[1].Value != "this is fatal" {
		t.Fatalf("expected \"this is fatal\" but got %q", rec.Log[1].Value)
	}
}
