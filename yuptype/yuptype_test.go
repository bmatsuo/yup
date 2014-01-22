// Copyright 2013, Bryan Matsuo. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// yuptype.go [created: Thu,  6 Jun 2013]

package yuptype

import (
	"fmt"
	"github.com/bmatsuo/yup"
	"github.com/bmatsuo/yup/yuptesting"
	"regexp"
	"testing"
)

func TestMsgSuffix(t*testing.T) {
	// tests the msgSuffix
	rec := yuptesting.Mock(func(t yuptesting.Test) { Equal(t, 1, 0, "a message") })
	match, err := regexp.MatchString(": expected 1 received 0; a message$", rec.Log[0].Value)
	yup.Assert(t, err == nil, "invalid regular expression")
	yup.Assert(t, match, fmt.Sprintf("unexpected error message content: %q", rec.Log[0].Value))
}

func TestEqual(t *testing.T) {
	rec := yuptesting.Mock(func(t yuptesting.Test) { Equal(t, "abc", "abc") })
	yup.T(t, 0, len(rec.Log) == 0, "unexpcted error")
	rec = yuptesting.Mock(func(t yuptesting.Test) { Equal(t, "abc", "def") })
	yup.T(t, 0, len(rec.Log) == 1, "expected 1 error but got", len(rec.Log))
	rec = yuptesting.Mock(func(t yuptesting.Test) { Equal(t, 0, int64(0)) })
	yup.T(t, 0, len(rec.Log) == 1, "expected 1 error but got", len(rec.Log))
	match, err := regexp.MatchString(": expected int[(]0[)] received int64[(]0[)]$", rec.Log[0].Value)
	yup.T(t, 0, err == nil, "invalid regular expression")
	yup.T(t, 0, match, fmt.Sprintf("unexpected error message content: %q", rec.Log[0].Value))
}

func TestNotEqual(t *testing.T) {
	rec := yuptesting.Mock(func(t yuptesting.Test) { NotEqual(t, "abc", "abc") })
	yup.T(t, 0, len(rec.Log) == 1, "expected 1 error but got", len(rec.Log))
	rec = yuptesting.Mock(func(t yuptesting.Test) { NotEqual(t, "abc", "def") })
	yup.T(t, 0, len(rec.Log) == 0, "unexpcted error")
}

func TestNil(t *testing.T) {
	rec := yuptesting.Mock(func(t yuptesting.Test) { Nil(t, nil) })
	yup.T(t, 0, len(rec.Log) == 0, "unexpected error")
	rec = yuptesting.Mock(func(t yuptesting.Test) { Nil(t, "abc") })
	yup.T(t, 0, len(rec.Log) == 1, "expected 1 error but got", len(rec.Log))
	var m map[string]interface{}
	rec = yuptesting.Mock(func(t yuptesting.Test) { Nil(t, m) })
	yup.T(t, 0, len(rec.Log) == 0, "unexpected error")
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
	rec := yuptesting.Mock(func(t yuptesting.Test) { Zero(t, nil) })
	Equal(t, 0, len(rec.Log), "unexpected number of errors")
	rec = yuptesting.Mock(func(t yuptesting.Test) { Zero(t, int64(0)) })
	Equal(t, 0, len(rec.Log), "unexpected number of errors")
	var empty []interface{}
	rec = yuptesting.Mock(func(t yuptesting.Test) { Zero(t, empty) })
	Equal(t, 0, len(rec.Log), "unexpected number of errors")

	rec = yuptesting.Mock(func(t yuptesting.Test) { Zero(t, make([]int, 0)) })
	Equal(t, 1, len(rec.Log), "unexpected number of errors")
	rec = yuptesting.Mock(func(t yuptesting.Test) { Zero(t, new(struct{})) })
	Equal(t, 1, len(rec.Log), "unexpected number of errors")
	rec = yuptesting.Mock(func(t yuptesting.Test) { Zero(t, "abc") })
	Equal(t, 1, len(rec.Log), "unexpected number of errors")
}

func TestNotZero(t *testing.T) {
	rec := yuptesting.Mock(func(t yuptesting.Test) { NotZero(t, nil) })
	Equal(t, 1, len(rec.Log), "unexpected number of errors")
	rec = yuptesting.Mock(func(t yuptesting.Test) { NotZero(t, int64(0)) })
	Equal(t, 1, len(rec.Log), "unexpected number of errors")
	var empty []interface{}
	rec = yuptesting.Mock(func(t yuptesting.Test) { NotZero(t, empty) })
	Equal(t, 1, len(rec.Log), "unexpected number of errors")

	rec = yuptesting.Mock(func(t yuptesting.Test) { NotZero(t, make([]int, 0)) })
	Equal(t, 0, len(rec.Log), "unexpected number of errors")
	rec = yuptesting.Mock(func(t yuptesting.Test) { NotZero(t, new(struct{})) })
	Equal(t, 0, len(rec.Log), "unexpected number of errors")
	rec = yuptesting.Mock(func(t yuptesting.Test) { NotZero(t, "abc") })
	Equal(t, 0, len(rec.Log), "unexpected number of errors")
}

func TestError(t *testing.T) {
	err := fmt.Errorf("")
	rec := yuptesting.Mock(func(t yuptesting.Test) { Error(t, err) })
	Equal(t, 0, len(rec.Log), "unexpected number of errors")
	rec = yuptesting.Mock(func(t yuptesting.Test) { Error(t, nil) })
	Equal(t, 1, len(rec.Log), "unexpected number of errors")
}
