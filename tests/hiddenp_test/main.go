package main

import (
	"github.com/01-edu/go-tests/lib/challenge"
	"github.com/01-edu/go-tests/lib/rand"
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
			[2]string{rand.RandLower()[:1], rand.RandLower()},
			[2]string{rand.RandUpper()[:1], rand.RandUpper()},
		)
	}
	for _, v := range args {
		challenge.Program("hiddenp", v[0], v[1])
	}
}
