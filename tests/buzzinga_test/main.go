package main

import (
    student "student"

    "github.com/01-edu/go-tests/lib/challenge"
    "github.com/01-edu/go-tests/solutions"
)

func main() {
    challenge.Function("BuzZinga", student.BuzZinga, solutions.BuzZinga, 20)
    challenge.Function("BuzZinga", student.BuzZinga, solutions.BuzZinga, 15)
    challenge.Function("BuzZinga", student.BuzZinga, solutions.BuzZinga, 10)
    challenge.Function("BuzZinga", student.BuzZinga, solutions.BuzZinga, 5)
    challenge.Function("BuzZinga", student.BuzZinga, solutions.BuzZinga, 4)
}
