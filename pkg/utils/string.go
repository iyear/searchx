package utils

import (
	"fmt"
	"github.com/thinkeridea/go-extend/exunicode/exutf8"
	"strconv"
	"time"
)

type str struct{}

var String = str{}

// Highlight start,end are valid string indexes. before,after are rune length
func (str str) Highlight(s string, start, end, before, after int, left, right string) string {
	if start >= end {
		return ""
	}

	s = s[:start] + left + s[start:end] + right + s[end:]

	start, end, before, after = start+len(left), end+len(right), before+len([]rune(left)), after+len([]rune(right))

	l, r, count := 0, 0, 0

	for index := range s {
		if index == start {
			l = count
		}

		if index == end {
			r = count
			break
		}

		count++
	}

	spacel, spacer := "...", "..."

	if l-before < 0 {
		l = before
		spacel = ""
	}

	if r+after >= len([]rune(s)) {
		spacer = ""
	}

	return spacel + exutf8.RuneSubString(s, l-before, (r+after)-(l-before)) + spacer

}

func (str str) MustGetDate(unix string) time.Time {
	u, err := strconv.ParseInt(unix, 10, 64)
	if err != nil {
		panic(fmt.Errorf("parse int failed: %s", unix))
	}

	return time.Unix(u, 0)
}

func (str str) RuneSubString(s string, l int) string {
	ss := exutf8.RuneSubString(s, 0, l)
	if len(ss) < len(s) {
		return ss + "..."
	}
	return ss
}
