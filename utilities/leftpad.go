package utilities

import (
	"fmt"
	"strings"
)

const (
	DefaultPadding = 4
	DefaultPadChar = " "
)

func LeftPad(s string, p int, c string) string {
	if p <= 0 {
		p = DefaultPadding
	}
	if c == "" {
		c = DefaultPadChar
	}
	return fmt.Sprintf("%s%s", strings.Repeat(c, p), s)
}
