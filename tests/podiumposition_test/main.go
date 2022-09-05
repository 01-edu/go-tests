package main

import (
	student "student"

	"github.com/01-edu/go-tests/lib/challenge"
	"github.com/01-edu/go-tests/solutions"
)

func main() {
	position := [][][]string {
	{{"4th Place"}, {"3rd Place"}, {"2nd Place"}, {"1st Place"}},
	{{"4th Place "}, {"3rd   Place"}, {"2nd Place  "}, {" 1st Place"}},
	{{""}, {"+55","*@Â£%D", "edr8927", "		| cat -e"}, {" "}},
	{{"741", "852","4","58","87","03","-96"},{"kd","jq","hel$lo","a","B","9W"}},
}
	for _, s := range position {
		challenge.Function("PodiumPosition", student.PodiumPosition, solutions.PodiumPosition, s)
	}
}
