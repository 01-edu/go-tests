package main

import (
	"github.com/01-edu/go-tests/lib/challenge"
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
	for i := 0; i < 30; i++ {
		args = append(args,
			[2]string{random.RandLower()[:1], random.RandLower()},
			[2]string{random.RandUpper()[:1], random.RandUpper()},
		)
	}
	for _, v := range args {
		challenge.Program("hiddenp", v[0], v[1])
	}
}
