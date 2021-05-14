package main

import (
	"github.com/01-edu/go-tests/lib/challenge"
	"github.com/01-edu/go-tests/lib/chars"
	"github.com/01-edu/go-tests/lib/random"
)

func main() {
	args := [][2]string{
		{"fgex.;", "tyf34gdgf;'ektufjhgdgex.;.;rtjynur6"},
		{"abc", "2altrb53c.sse"},
		{"abc", "btarc"},
		{"DD", "DABC"},
		{""},
	}
	for i := 0; i < 5; i++ {
		args = append(args,
			[2]string{random.Str(chars.Lower, 1), random.Str(chars.Lower, 13)},
			[2]string{random.Str(chars.Upper, 1), random.Str(chars.Upper, 13)},
		)
	}
	for _, v := range args {
		challenge.Program("hiddenp", v[0], v[1])
	}
}
