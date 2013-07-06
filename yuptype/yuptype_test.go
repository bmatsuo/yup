// Copyright 2013, Bryan Matsuo. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// yuptype.go [created: Thu,  6 Jun 2013]

package yuptype

import (
	"github.com/bmatsuo/yup"
	"github.com/bmatsuo/yup/yuptesting"
	"testing"
)

func TestEqual(t *testing.T) {
	rec := yuptesting.Mock(func(t yuptesting.Test) { Equal(t, "abc", "abc") })
	yup.Assert(t, 0, len(rec.Log) == 0, "unexpcted error")
	rec = yuptesting.Mock(func(t yuptesting.Test) { Equal(t, "abc", "def") })
	yup.Assert(t, 0, len(rec.Log) == 1, "expected 1 error but got", len(rec.Log))
}

func TestNotEqual(t *testing.T) {
	rec := yuptesting.Mock(func(t yuptesting.Test) { NotEqual(t, "abc", "abc") })
	yup.Assert(t, 0, len(rec.Log) == 1, "expected 1 error but got", len(rec.Log))
	rec = yuptesting.Mock(func(t yuptesting.Test) { NotEqual(t, "abc", "def") })
	yup.Assert(t, 0, len(rec.Log) == 0, "unexpcted error")
}

func TestNil(t *testing.T) {
	rec := yuptesting.Mock(func(t yuptesting.Test) { Nil(t, nil) })
	yup.Assert(t, 0, len(rec.Log) == 0, "unexpected error")
	rec = yuptesting.Mock(func(t yuptesting.Test) { Nil(t, "abc") })
	yup.Assert(t, 0, len(rec.Log) == 1, "expected 1 error but got", len(rec.Log))
	var m map[string]interface{}
	rec = yuptesting.Mock(func(t yuptesting.Test) { Nil(t, m) })
	yup.Assert(t, 0, len(rec.Log) == 0, "unexpected error")
}

func TestNotNil(t *testing.T) {
	rec := yuptesting.Mock(func(t yuptesting.Test) { NotNil(t, nil) })
	Equal(t, 1, len(rec.Log), "unexpected number of errors")
	rec = yuptesting.Mock(func(t yuptesting.Test) { NotNil(t, "abc") })
	Equal(t, 0, len(rec.Log), "unexpected number of errors")
	var m map[string]interface{}
	rec = yuptesting.Mock(func(t yuptesting.Test) { NotNil(t, m) })
	Equal(t, 1, len(rec.Log), "unexpected number of errors")
}

func TestZero(t *testing.T) {
	t.Fatal("unimplemented")
}

func TestNotZero(t *testing.T) {
	t.Fatal("unimplemented")
}

func TestError(t *testing.T) {
	t.Fatal("unimplemented")
}

func TestNotError(t *testing.T) {
	t.Fatal("unimplemented")
}
