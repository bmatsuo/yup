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

func msgSuffix(msg []interface{}) string {
	if len(msg) == 0 {
		return ""
	}
	return "; " + fmt.Sprint(msg...)
}

// deep structural equality between expected and val.
func Equal(t yup.Test, expected, val interface{}, msg ...interface{}) {
	yup.TD(t, 1, reflect.DeepEqual(expected, val), func() string {
		// messy logic to deal with hard-to-spot corner cases.
		exstr := fmt.Sprintf("%#v", expected)
		vstr := fmt.Sprintf("%#v", val)

		if exstr == vstr {
			// naively attempt to differenciate similar looking things.
			_exstr := fmt.Sprintf("%T(%s)", expected, exstr)
			_vstr := fmt.Sprintf("%T(%s)", val, vstr)
			if _exstr != _vstr {
				exstr, vstr = _exstr, _vstr
			}
		}

		return fmt.Sprintf("expected %s received %s%s", exstr, vstr, msgSuffix(msg))
	})
}

// the opposite of Equal().
func NotEqual(t yup.Test, unexpected, val interface{}, msg ...interface{}) {
	yup.T(t, 1, !reflect.DeepEqual(unexpected, val),
		fmt.Sprintf("unexpected %s%v",
			fmt.Sprintf("%#v", val),
			msgSuffix(msg)))
}

// val is nil.
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
	yup.T(t, 1, isNil, fmt.Sprintf("unexpected non-nil value (%v)%v", val, msgSuffix(msg)))
}

// the opposite of Nil().
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
	yup.T(t, 1, !isNil, "unexpected nil value"+msgSuffix(msg))
}

// val is the zero value of it's type.
func Zero(t yup.Test, val interface{}, msg ...interface{}) {
	yup.T(t, 1, isZero(val),
		fmt.Sprintf("unexpected non-zero value (%#v)%v",
			val, msgSuffix(msg)))
}

// the opposite of Zero().
func NotZero(t yup.Test, val interface{}, msg ...interface{}) {
	yup.T(t, 1, !isZero(val),
		fmt.Sprintf("unexpected zero value (%#v)%v",
			val, msgSuffix(msg)))
}

func isZero(val interface{}) bool {
	if val == nil {
		return true
	}
	typ := reflect.TypeOf(val)
	zero := reflect.Zero(typ)
	return reflect.DeepEqual(val, zero.Interface())
}

// err is not nil.
func Error(t yup.Test, err error, msg ...interface{}) {
	yup.T(t, 1, err != nil, "expected error"+msgSuffix(msg))
}
