// Copyright 2013, Bryan Matsuo. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// yaaptype.go [created: Thu,  6 Jun 2013]

// Package yaaptype does ....
package yaaptype

import (
	"github.com/bmatsuo/yaap"

	"fmt"
	"reflect"
)

func Equal(t yaap.Test, expected, val interface{}, msg ...interface{}) {
	yaap.Assert(t, 1, reflect.DeepEqual(expected, val),
		fmt.Sprintf("expected %#v but received %#v", expected, val),
		fmt.Sprint(msg...))
}

func NotEqual(t yaap.Test, expected, val interface{}, msg ...interface{}) {
	yaap.Assert(t, 1, !reflect.DeepEqual(expected, val),
		fmt.Sprintf("expected %#v but received %#v", expected, val),
		fmt.Sprint(msg...))
}

// is this ever different than Zero() for pointer types?
func Nil(t yaap.Test, val interface{}, msg ...interface{}) {
	yaap.Assert(t, 1, val == nil, "unexpected nil value; ", fmt.Sprint(msg...))
}

// is this ever different than NotZero() for pointer types?
func NotNil(t yaap.Test, val interface{}, msg ...interface{}) {
	yaap.Assert(t, 1, val != nil, "unexpected nil value; ", fmt.Sprint(msg...))
}

func True(t yaap.Test, ok bool, msg ...interface{}) {
	yaap.Assert(t, 1, ok, "unexpected false value; ", fmt.Sprint(msg...))
}

func False(t yaap.Test, ok bool, msg ...interface{}) {
	yaap.Assert(t, 1, !ok, "unexpected true value; ", fmt.Sprint(msg...))
}

func Zero(t yaap.Test, val interface{}, msg ...interface{}) {
	yaap.Assert(t, 1, isZero(val), "unexpected non-zero value; ", fmt.Sprint(msg...))
}

func NotZero(t yaap.Test, val interface{}, msg ...interface{}) {
	yaap.Assert(t, 1, !isZero(val), "unexpected zero value; ", fmt.Sprint(msg...))
}

func isZero(val interface{}) bool {
	typ := reflect.TypeOf(val)
	zero := reflect.Zero(typ)
	return reflect.DeepEqual(val, zero.Interface())
}
