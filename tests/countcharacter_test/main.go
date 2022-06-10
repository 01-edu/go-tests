package main

import (
	student "student"

	"github.com/01-edu/go-tests/lib/base"
	"github.com/01-edu/go-tests/lib/challenge"
	"github.com/01-edu/go-tests/solutions"
)

func CountChar(str string, c rune) int {
	var count int
	for _, v := range str {
		if rune(v) == c {
			count++
		}
	}
	return count
}

func main() {
	challenge.Function("CountChar", student.CountChar, CountChar, "hello", 'e')
	challenge.Function("CountChar", student.CountChar, CountChar, " 	", ' ')
	challenge.Function("CountChar", student.CountChar, CountChar, "the 7 deadly sins", '7')
}
