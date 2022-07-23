package main

import (
	student "student"

	"github.com/01-edu/go-tests/lib/challenge"
	"github.com/01-edu/go-tests/solutions"
)

func main() {
	challenge.Function("CountChar", solutions.CountChar, student.CountChar, "hello", 'e')
	challenge.Function("CountChar", solutions.CountChar, student.CountChar, "	 	", ' ')
	challenge.Function("CountChar", solutions.CountChar, student.CountChar, "the 7 deadly sins", '7')
	challenge.Function("CountChar", solutions.CountChar, student.CountChar, "the 7 deadly sins", 's')
	challenge.Function("CountChar", solutions.CountChar, student.CountChar, "the 7 deadly sins", 'd')
	challenge.Function("CountChar", solutions.CountChar, student.CountChar, "", 'i')
}
