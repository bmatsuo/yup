// Copyright 2013, Bryan Matsuo. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// yaaptype.go [created: Thu,  6 Jun 2013]

package yaaptype

import (
	"github.com/bmatsuo/yaap"
	"github.com/bmatsuo/yaap/yaaptesting"
	"testing"
)

func TestEqual(t *testing.T) {
	rec := yaaptesting.Mock(func(t yaaptesting.Test) { Equal(t, "abc", "abc") })
	yaap.Assert(t, 0, len(rec.Log) == 0, "unexpcted error")
	rec = yaaptesting.Mock(func(t yaaptesting.Test) { Equal(t, "abc", "def") })
	yaap.Assert(t, 0, len(rec.Log) == 1, "expected 1 error but got", len(rec.Log))
}

func TestNotEqual(t *testing.T) {
	rec := yaaptesting.Mock(func(t yaaptesting.Test) { NotEqual(t, "abc", "abc") })
	yaap.Assert(t, 0, len(rec.Log) == 1, "expected 1 error but got", len(rec.Log))
	rec = yaaptesting.Mock(func(t yaaptesting.Test) { NotEqual(t, "abc", "def") })
	yaap.Assert(t, 0, len(rec.Log) == 0, "unexpcted error")
}

func TestNil(t *testing.T) {
	rec := yaaptesting.Mock(func(t yaaptesting.Test) { Nil(t, nil) })
	yaap.Assert(t, 0, len(rec.Log) == 0, "unexpected error")
	rec = yaaptesting.Mock(func(t yaaptesting.Test) { Nil(t, "abc", "def") })
	yaap.Assert(t, 0, len(rec.Log) == 1, "expected 1 error but got", len(rec.Log))
}
