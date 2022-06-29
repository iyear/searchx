package util

import (
	"fmt"
	"github.com/thinkeridea/go-extend/exunicode/exutf8"
	"strconv"
	"time"
)

//Highlight start,end are valid string indexes. before,after are rune length
func Highlight(s string, start, end, before, after int, highlight string) string {
	if start >= end {
		return ""
	}

	s = s[:start] + highlight + s[start:end] + highlight + s[end:]

	start, end, before, after = start+len(highlight), end+len(highlight), before+len([]rune(highlight)), after+len([]rune(highlight))

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

	if l-before < 0 {
		l = before
	}

	return exutf8.RuneSubString(s, l-before, (r+after)-(l-before))
}

func MustGetDate(unix string) time.Time {
	u, err := strconv.ParseInt(unix, 10, 64)
	if err != nil {
		panic(fmt.Errorf("parse int failed: %s", unix))
	}

	return time.Unix(u, 0)
}
