package utilities

import (
	"fmt"
	"strings"
)

const (
	DefaultRightPadding = 4
	DefaultRightPadChar = " "
)

func RightPad(s string, p int, c string) string {
	if p <= 0 {
		p = DefaultRightPadding
	}
	if c == "" {
		c = DefaultRightPadChar
	}
	return fmt.Sprintf("%s%s", s, strings.Repeat(c, p))
}
