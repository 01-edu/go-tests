package chars

import (
	"strings"
	"unicode/utf8"
)

var (
	Digit = runeRange('0', '9')             // Decimal digit characters
	Lower = runeRange('a', 'z')             // Lowercase latin alphabet characters
	Upper = runeRange('A', 'Z')             // Uppercase latin alphabet characters
	ASCII = runeRange(' ', '~')             // ASCII printable characters
	Basic = Lower + Upper                   // Lower and Upper characters
	Alnum = Basic + Digit                   // Basic and Digit characters
	Words = Alnum + strings.Repeat(" ", 10) // Alnum and Space characters
)

func runeRange(a, b rune) string {
	var s []rune
	for {
		if utf8.ValidRune(a) {
			s = append(s, a)
		}
		if a == b {
			return string(s)
		}
		if a < b {
			a++
		} else {
			a--
		}
	}
}
