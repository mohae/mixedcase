// Copyright (c) 2016, Joel Scoble. All rights reserved.
//
// Use of this source code is governed by the MIT license
// that can be found in the LICENSE file.
package mixedcase

import (
	"fmt"
	"strings"
	"unicode"
	"unicode/utf8"
)

// Exported returns the string as a MixedCase string that can be used for
// exported identifiers.
func Exported(s string) string {
	var v string

	ndx := discardStart(s)
	if ndx > 0 {
		s = s[ndx:]
	}
	vals := strings.Split(s, "_")
	for i, val := range vals {
		if i == 0 {
			val = NumToAlpha(val)
		}
		v = fmt.Sprintf("%s%s", v, UpperInitialism(strings.Title(val)))
	}
	return v
}

// Unexported returns the string as a mixedCase string that can be used for
// unexported identifiers.
func Unexported(s string) string {
	var v string

	ndx := discardStart(s)
	if ndx > 0 {
		s = s[ndx:]
	}
	vals := strings.Split(s, "_")
	for i, val := range vals {
		if i == 0 {
			v = NumToAlpha(val)
			v = LowerInitialism(LowerFirstRune(v))
			continue
		}
		v = fmt.Sprintf("%s%s", v, UpperInitialism(strings.Title(val)))
	}
	return v
}

// List and comment is from https://github.com/golang/lint/blob/master/lint.go
// Original copyright:
// Copyright (c) 2013 The Go Authors. All rights reserved.
//
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file or at
// https://developers.google.com/open-source/licenses/bsd.
//
// commonInitialisms is a set of common initialisms.
// Only add entries that are highly unlikely to be non-initialisms.
// For instance, "ID" is fine (Freudian code is rare), but "AND" is not.
var commonInitialisms = map[string]struct{}{
	"API":   struct{}{},
	"ASCII": struct{}{},
	"CPU":   struct{}{},
	"CSS":   struct{}{},
	"DNS":   struct{}{},
	"EOF":   struct{}{},
	"GUID":  struct{}{},
	"HTML":  struct{}{},
	"HTTP":  struct{}{},
	"HTTPS": struct{}{},
	"ID":    struct{}{},
	"IP":    struct{}{},
	"JSON":  struct{}{},
	"LHS":   struct{}{},
	"QPS":   struct{}{},
	"RAM":   struct{}{},
	"RHS":   struct{}{},
	"RPC":   struct{}{},
	"SLA":   struct{}{},
	"SMTP":  struct{}{},
	"SNI":   struct{}{},
	"SSH":   struct{}{},
	"TLS":   struct{}{},
	"TTL":   struct{}{},
	"UI":    struct{}{},
	"UID":   struct{}{},
	"UUID":  struct{}{},
	"URI":   struct{}{},
	"URL":   struct{}{},
	"UTF8":  struct{}{},
	"VM":    struct{}{},
	"XML":   struct{}{},
}

// UpperInitialism returns the string as all UPPER case if it matches a
// supported initialism, otherwise the original value is returned.
func UpperInitialism(s string) string {
	tmp := strings.ToUpper(s)
	if _, ok := commonInitialisms[tmp]; ok {
		return tmp
	}
	return s
}

// LowerInitialism returns the string as all lower case if it matches a
// supported initialism, otherwise the original value is returned.
func LowerInitialism(s string) string {
	tmp := strings.ToUpper(s)
	if _, ok := commonInitialisms[tmp]; ok {
		return strings.ToLower(tmp)
	}
	return s
}

// discardStart checks the beginning of the string for characters that should
// be discarded until it finds a value that should not be discarded.
func discardStart(s string) int {
	var pos int
	for i, w := 0, 0; i < len(s); i += w {
		v, width := utf8.DecodeRuneInString(s[i:])
		w = width
		if shouldDiscard(v) {
			continue
		}
		pos = i
		break
	}
	return pos
}

func shouldDiscard(r rune) bool {
	switch r {
	case '~', '!', '@', '#', '$', '%', '^', '&', '*', '-', '_', '=', '+', ':', '.', '<', '>':
		return true
	}
	return false
}

// NumToAlpha checks to see if the first char is a number, if it is, it gets
// converted to its word equivalent. The rest of the string is uppercased if
// it is an initialism and title cased since it now counts as a seperate word.
func NumToAlpha(s string) string {
	var n string
	switch s[0] {
	case '0':
		n = "Zero"
	case '1':
		n = "One"
	case '2':
		n = "Two"
	case '3':
		n = "Three"
	case '4':
		n = "Four"
	case '5':
		n = "Five"
	case '6':
		n = "Six"
	case '7':
		n = "Seven"
	case '8':
		n = "Eight"
	case '9':
		n = "Nine"
	}
	if n == "" {
		return s
	}
	return fmt.Sprintf("%s%s", n, UpperInitialism(strings.Title(s[1:])))
}

// LowerFirstRune returns a copy of the string s with the first Unicode letter
// mapped to its lowercase.
func LowerFirstRune(s string) string {
	if len(s) == 0 {
		return s
	}
	r, w := utf8.DecodeRuneInString(s[0:])
	r = unicode.ToLower(r)
	return fmt.Sprintf("%c%s", r, s[w:])
}
