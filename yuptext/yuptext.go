// Copyright 2013, Bryan Matsuo. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// yuptext.go [created: Sat,  6 Jul 2013]

/*
yuptext helps make assertions on text.
*/
package yuptext

import (
	"github.com/bmatsuo/yup"

	"bytes"
	"fmt"
	"regexp"
	"strings"
)

// show at least n characters s.
func trunc(s string, n int) string {
	if n <= 0 {
		return s
	}
	if len(s) <= n + 3 {
		return s
	}
	return s[:n] + "..."
}

func HasPrefix(t yup.Test, p, prefix []byte, msg ...interface{}) {
	yup.T(t, 1, bytes.HasPrefix(p, prefix),
		fmt.Sprintf("expected prefix %q on %q; %v",
			prefix, trunc(string(p), len(prefix)+3), fmt.Sprint(msg...)))
}

func HasPrefixString(t yup.Test, str, prefix string, msg ...interface{}) {
	yup.T(t, 1, strings.HasPrefix(str, prefix),
		fmt.Sprintf("expected prefix %q on %q; %v",
			prefix, trunc(str, len(prefix)+3), fmt.Sprint(msg...)))
}

func HasSuffix(t yup.Test, p, suffix []byte, msg ...interface{}) {
	yup.T(t, 1, bytes.HasSuffix(p, suffix),
		fmt.Sprintf("expected suffix %q on %q; %v",
			suffix, string(p), fmt.Sprint(msg...)))
}

func HasSuffixString(t yup.Test, str, suffix string, msg ...interface{}) {
	yup.T(t, 1, strings.HasSuffix(str, suffix),
		fmt.Sprintf("expected suffix %q on %q; %v",
			suffix, str, fmt.Sprint(msg...)))
}

func Contains(t yup.Test, p, segment []byte, msg ...interface{}) {
	yup.T(t, 1, bytes.Contains(p, segment),
		fmt.Sprintf("expected segment %q in %q; %v",
			segment, string(p), fmt.Sprint(msg...)))
}

func ContainsString(t yup.Test, str, substr string, msg ...interface{}) {
	yup.T(t, 1, strings.Contains(str, substr),
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
	yup.T(t, 1, r.Match(p),
		fmt.Sprintf("pattern %q does not match %q; %v",
			r.String(), string(p), fmt.Sprint(msg...)))
}

func MatchPattString(t yup.Test, str, patt string, msg ...interface{}) {
	r := compileRegexp(t, 1, patt, msg...)
	yup.T(t, 1, r.MatchString(str),
		fmt.Sprintf("pattern %q does not match %q; %v",
			r.String(), str, fmt.Sprint(msg...)))
}

func compileRegexp(t yup.Test, n uint, patt string, msg ...interface{}) *regexp.Regexp {
	r, err := regexp.Compile(patt)
	yup.T(t, 1+n, err == nil,
		fmt.Sprintf("invalid regular expression %q; %v",
			patt, fmt.Sprint(msg...)))
	return r
}
