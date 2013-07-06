// Copyright 2013, Bryan Matsuo. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// yuptext_test.go [created: Sat,  6 Jul 2013]

package yuptext

import (
	"testing"

	"github.com/bmatsuo/yup/yuptesting"
	"github.com/bmatsuo/yup/yuptype"
)

func TestHasPrefix(t *testing.T) {
	p := []byte{0, 1, 2}
	empty := make([]byte, 0)

	errlog := yuptesting.Mock(func(t yuptesting.Test) { HasPrefix(t, p, nil) })
	yuptype.Equal(t, 0, len(errlog.Log))
	errlog = yuptesting.Mock(func(t yuptesting.Test) { HasPrefix(t, p, p[:1]) })
	yuptype.Equal(t, 0, len(errlog.Log))
	errlog = yuptesting.Mock(func(t yuptesting.Test) { HasPrefix(t, nil, nil) })
	yuptype.Equal(t, 0, len(errlog.Log))
	errlog = yuptesting.Mock(func(t yuptesting.Test) {
		HasPrefix(t, nil, empty)
	})
	yuptype.Equal(t, 0, len(errlog.Log))
	errlog = yuptesting.Mock(func(t yuptesting.Test) {
		HasPrefix(t, empty, nil)
	})
	yuptype.Equal(t, 0, len(errlog.Log))

	errlog = yuptesting.Mock(func(t yuptesting.Test) { HasPrefix(t, nil, p) })
	yuptype.Equal(t, 1, len(errlog.Log))
	errlog = yuptesting.Mock(func(t yuptesting.Test) { HasPrefix(t, p, p[1:]) })
	yuptype.Equal(t, 1, len(errlog.Log))
}

func TestHasPrefixString(t *testing.T) {
	str := "012"

	errlog := yuptesting.Mock(func(t yuptesting.Test) {
		HasPrefixString(t, str, "")
	})
	yuptype.Equal(t, 0, len(errlog.Log))
	errlog = yuptesting.Mock(func(t yuptesting.Test) {
		HasPrefixString(t, str, str[:1])
	})
	yuptype.Equal(t, 0, len(errlog.Log))
	errlog = yuptesting.Mock(func(t yuptesting.Test) {
		HasPrefixString(t, "", "")
	})
	yuptype.Equal(t, 0, len(errlog.Log))

	errlog = yuptesting.Mock(func(t yuptesting.Test) {
		HasPrefixString(t, "", str)
	})
	yuptype.Equal(t, 1, len(errlog.Log))
	errlog = yuptesting.Mock(func(t yuptesting.Test) {
		HasPrefixString(t, str, str[1:])
	})
	yuptype.Equal(t, 1, len(errlog.Log))
}

func TestHasSuffix(t *testing.T) {
	p := []byte{0, 1, 2}
	empty := make([]byte, 0)

	errlog := yuptesting.Mock(func(t yuptesting.Test) { HasSuffix(t, p, nil) })
	yuptype.Equal(t, 0, len(errlog.Log))
	errlog = yuptesting.Mock(func(t yuptesting.Test) { HasSuffix(t, p, p[1:]) })
	yuptype.Equal(t, 0, len(errlog.Log))
	errlog = yuptesting.Mock(func(t yuptesting.Test) { HasSuffix(t, nil, nil) })
	yuptype.Equal(t, 0, len(errlog.Log))
	errlog = yuptesting.Mock(func(t yuptesting.Test) {
		HasSuffix(t, nil, empty)
	})
	yuptype.Equal(t, 0, len(errlog.Log))
	errlog = yuptesting.Mock(func(t yuptesting.Test) {
		HasSuffix(t, empty, nil)
	})
	yuptype.Equal(t, 0, len(errlog.Log))

	errlog = yuptesting.Mock(func(t yuptesting.Test) { HasSuffix(t, nil, p) })
	yuptype.Equal(t, 1, len(errlog.Log))
	errlog = yuptesting.Mock(func(t yuptesting.Test) { HasSuffix(t, p, p[:1]) })
	yuptype.Equal(t, 1, len(errlog.Log))
}

func TestHasSuffixString(t *testing.T) {
	str := "012"

	errlog := yuptesting.Mock(func(t yuptesting.Test) {
		HasSuffixString(t, str, "")
	})
	yuptype.Equal(t, 0, len(errlog.Log))
	errlog = yuptesting.Mock(func(t yuptesting.Test) {
		HasSuffixString(t, str, str[1:])
	})
	yuptype.Equal(t, 0, len(errlog.Log))
	errlog = yuptesting.Mock(func(t yuptesting.Test) {
		HasSuffixString(t, "", "")
	})
	yuptype.Equal(t, 0, len(errlog.Log))

	errlog = yuptesting.Mock(func(t yuptesting.Test) {
		HasSuffixString(t, "", str)
	})
	yuptype.Equal(t, 1, len(errlog.Log))
	errlog = yuptesting.Mock(func(t yuptesting.Test) {
		HasSuffixString(t, str, str[:1])
	})
	yuptype.Equal(t, 1, len(errlog.Log))
}

func TestContains(t *testing.T) {
	p := []byte{0, 1, 2}
	empty := make([]byte, 0)

	errlog := yuptesting.Mock(func(t yuptesting.Test) { Contains(t, p, nil) })
	yuptype.Equal(t, 0, len(errlog.Log))
	errlog = yuptesting.Mock(func(t yuptesting.Test) { Contains(t, p, p[:1]) })
	yuptype.Equal(t, 0, len(errlog.Log))
	errlog = yuptesting.Mock(func(t yuptesting.Test) { Contains(t, p, p[1:]) })
	yuptype.Equal(t, 0, len(errlog.Log))
	errlog = yuptesting.Mock(func(t yuptesting.Test) { Contains(t, p, p) })
	yuptype.Equal(t, 0, len(errlog.Log))
	errlog = yuptesting.Mock(func(t yuptesting.Test) { Contains(t, nil, nil) })
	yuptype.Equal(t, 0, len(errlog.Log))
	errlog = yuptesting.Mock(func(t yuptesting.Test) {
		Contains(t, nil, empty)
	})
	yuptype.Equal(t, 0, len(errlog.Log))
	errlog = yuptesting.Mock(func(t yuptesting.Test) {
		Contains(t, empty, nil)
	})
	yuptype.Equal(t, 0, len(errlog.Log))

	errlog = yuptesting.Mock(func(t yuptesting.Test) { Contains(t, nil, p) })
	yuptype.Equal(t, 1, len(errlog.Log))
	superp := append(p, 1)
	errlog = yuptesting.Mock(func(t yuptesting.Test) { Contains(t, p, superp) })
	yuptype.Equal(t, 1, len(errlog.Log))
	diffp := append([]byte{}, p...)
	diffp[1] = p[1] + 1
	errlog = yuptesting.Mock(func(t yuptesting.Test) { Contains(t, p, diffp) })
	yuptype.Equal(t, 1, len(errlog.Log))
}

func TestContainsString(t *testing.T) {
	str := "012"

	errlog := yuptesting.Mock(func(t yuptesting.Test) {
		ContainsString(t, str, "")
	})
	yuptype.Equal(t, 0, len(errlog.Log))
	errlog = yuptesting.Mock(func(t yuptesting.Test) {
		ContainsString(t, str, str[:1])
	})
	yuptype.Equal(t, 0, len(errlog.Log))
	errlog = yuptesting.Mock(func(t yuptesting.Test) {
		ContainsString(t, str, str[1:])
	})
	yuptype.Equal(t, 0, len(errlog.Log))
	errlog = yuptesting.Mock(func(t yuptesting.Test) {
		ContainsString(t, str, str)
	})
	yuptype.Equal(t, 0, len(errlog.Log))
	errlog = yuptesting.Mock(func(t yuptesting.Test) {
		ContainsString(t, "", "")
	})
	yuptype.Equal(t, 0, len(errlog.Log))

	errlog = yuptesting.Mock(func(t yuptesting.Test) {
		ContainsString(t, "", str)
	})
	yuptype.Equal(t, 1, len(errlog.Log))
	errlog = yuptesting.Mock(func(t yuptesting.Test) {
		ContainsString(t, str, str + "1")
	})
	yuptype.Equal(t, 1, len(errlog.Log))
	diffstr := str[:1] + "x" + str[2:]
	errlog = yuptesting.Mock(func(t yuptesting.Test) {
		ContainsString(t, str, diffstr)
	})
	yuptype.Equal(t, 1, len(errlog.Log))
}
