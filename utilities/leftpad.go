package utilities

import (
	"fmt"
	"strings"
)

const (
	DefaultLeftPadding = 4
	DefaultLeftPadChar = " "
)

func LeftPad(s string, p int, c string) string {
	if p <= 0 {
		p = DefaultLeftPadding
	}
	if c == "" {
		c = DefaultLeftPadChar
	}
	return fmt.Sprintf("%s%s", strings.Repeat(c, p), s)
}
