// Copyright 2013, Bryan Matsuo. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// yupstring.go [created: Sat,  6 Jul 2013]

/*
Yupstring helps make assertions on text.
*/
package yuptext

import (
	"github.com/bmatsuo/yup"

	"bytes"
	"fmt"
	"regexp"
	"strings"
)

var TruncLen int

func trunc(s string) string {
	if TruncLen <= 0 {
		return s
	}
	if len(s) <= TruncLen {
		return s
	}
	return s[:TruncLen-3] + "..."
}

func HasPrefix(t yup.Test, p, prefix []byte, msg ...interface{}) {
	yup.Assert(t, 1, bytes.HasPrefix(p, prefix),
		fmt.Sprintf("expected prefix %q on %q; %v",
			prefix, trunc(string(p)), fmt.Sprint(msg...)))
}

func HasPrefixString(t yup.Test, str, prefix string, msg ...interface{}) {
	yup.Assert(t, 1, strings.HasPrefix(str, prefix),
		fmt.Sprintf("expected prefix %q on %q; %v",
			prefix, trunc(str), fmt.Sprint(msg...)))
}

func HasSuffix(t yup.Test, p, suffix []byte, msg ...interface{}) {
	yup.Assert(t, 1, bytes.HasSuffix(p, suffix),
		fmt.Sprintf("expected suffix %q on %q; %v",
			suffix, trunc(string(p)), fmt.Sprint(msg...)))
}

func HasSuffixString(t yup.Test, str, suffix string, msg ...interface{}) {
	yup.Assert(t, 1, strings.HasSuffix(str, suffix),
		fmt.Sprintf("expected suffix %q on %q; %v",
			suffix, trunc(str), fmt.Sprint(msg...)))
}

func Contains(t yup.Test, p, segment []byte, msg ...interface{}) {
	yup.Assert(t, 1, bytes.Contains(p, segment),
		fmt.Sprintf("expected segment %q in %q; %v",
			segment, trunc(string(p)), fmt.Sprint(msg...)))
}

func ContainsString(t yup.Test, str, substr string, msg ...interface{}) {
	yup.Assert(t, 1, strings.Contains(str, substr),
		fmt.Sprintf("expected substring %q in %q; %v",
			substr, str, fmt.Sprint(msg...)))
}

/*
TODO these need an assertion that r is not nil
func Match(t yup.Test, p []byte, r *regexp.Regexp, msg ...interface{}) {
	matchRegexp(t, 1, string(p), r, msg...)
}

func MatchString(t yup.Test, str string, r *regexp.Regexp, msg ...interface{}) {
	matchRegexp(t, 1, str, r, msg...)
}
*/

func MatchPatt(t yup.Test, p []byte, patt string, msg ...interface{}) {
	r := compileRegexp(t, 1, patt, msg...)
	matchRegexp(t, 1, string(p), r, msg...)
}

func MatchPattString(t yup.Test, str, patt string, msg ...interface{}) {
	r := compileRegexp(t, 1, patt, msg...)
	matchRegexp(t, 1, str, r, msg...)
}

func compileRegexp(t yup.Test, n int, patt string, msg ...interface{}) *regexp.Regexp {
	r, err := regexp.Compile(patt)
	yup.Assert(t, 1+n, err == nil,
		fmt.Sprintf("invalid regular expression %q; %v",
			patt, fmt.Sprint(msg...)))
	return r
}

func matchRegexp(t yup.Test, n int, str string, r *regexp.Regexp, msg ...interface{}) {
	yup.Assert(t, 1+n, r.MatchString(str),
		fmt.Sprintf("pattern %q does not match %q; %v",
			r.String(), trunc(str), fmt.Sprint(msg...)))
}
