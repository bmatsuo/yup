// Copyright 2013, Bryan Matsuo. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// yuptype.go [created: Thu,  6 Jun 2013]

/*
Implements generic type-based assertions, generally using reflection.
*/
package yuptype

import (
	"github.com/bmatsuo/yup"

	"fmt"
	"reflect"
)

// Deep structural equality between expected and val.
func Equal(t yup.Test, expected, val interface{}, msg ...interface{}) {
	yup.Assert(t, 1, reflect.DeepEqual(expected, val),
		fmt.Sprintf("expected %#v but received %#v; %v",
			expected, val, fmt.Sprint(msg...)))
}

// The opposite of Equal().
func NotEqual(t yup.Test, expected, val interface{}, msg ...interface{}) {
	yup.Assert(t, 1, !reflect.DeepEqual(expected, val),
		fmt.Sprintf("expected %#v but received %#v; %v",
			expected, val, fmt.Sprint(msg...)))
}

// Val is nil.
// is this ever different than Zero() for pointer types?
func Nil(t yup.Test, val interface{}, msg ...interface{}) {
	v := reflect.ValueOf(val)
	isNil := !v.IsValid()
	switch v.Kind() {
	case reflect.Slice,
		reflect.Chan,
		reflect.Map,
		reflect.Ptr,
		reflect.Interface,
		reflect.Func:
		isNil = v.IsNil()
	}
	yup.Assert(t, 1, isNil, "unexpected non-nil value; ", fmt.Sprint(msg...))
}

// The opposite of Nil().
// is this ever different than NotZero() for pointer types?
func NotNil(t yup.Test, val interface{}, msg ...interface{}) {
	v := reflect.ValueOf(val)
	isNil := !v.IsValid()
	switch v.Kind() {
	case reflect.Slice,
		reflect.Chan,
		reflect.Map,
		reflect.Ptr,
		reflect.Interface,
		reflect.Func:
		isNil = v.IsNil()
	}
	yup.Assert(t, 1, !isNil, "unexpected nil value; ", fmt.Sprint(msg...))
}

// Val is the zero value of it's type.
func Zero(t yup.Test, val interface{}, msg ...interface{}) {
	yup.Assert(t, 1, isZero(val),
		fmt.Sprintf("unexpected non-zero value (%#v); %v",
			val, fmt.Sprint(msg...)))
}

// The opposite of Zero().
func NotZero(t yup.Test, val interface{}, msg ...interface{}) {
	yup.Assert(t, 1, !isZero(val),
		fmt.Sprintf("unexpected zero value (%#v); %v",
			val, fmt.Sprint(msg...)))
}

func isZero(val interface{}) bool {
	if val == nil {
		return true
	}
	typ := reflect.TypeOf(val)
	zero := reflect.Zero(typ)
	return reflect.DeepEqual(val, zero.Interface())
}

// Err is not nil.
func Error(t yup.Test, err error, msg ...interface{}) {
	yup.Assert(t, 1, err != nil, "expected error; ", fmt.Sprint(msg...))
}

// The opposite of Error().
func NotError(t yup.Test, err error, msg ...interface{}) {
	yup.Assert(t, 1, err == nil,
		fmt.Sprintf("unexpected error (%v); %v",
			err, fmt.Sprint(msg...)))
}

// Ok is true.
func True(t yup.Test, ok bool, msg ...interface{}) {
	yup.Assert(t, 1, ok, "unexpected false value; ", fmt.Sprint(msg...))
}

// Ok is false.
func False(t yup.Test, ok bool, msg ...interface{}) {
	yup.Assert(t, 1, !ok, "unexpected true value; ", fmt.Sprint(msg...))
}
