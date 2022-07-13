package main

import (
	"strings"

	"github.com/01-edu/go-tests/lib/challenge"
)

func main() {
	table := []string{
		"1", "2", "3", "4", "5", "6", "7", "8", "9", "10",
		"11", "12", "13", "14", "15", "16", "17", "18", "19", "20",
		"01", "-1", "-2", "-3", "-4", "-5", "-6", "-7", "-8", "-9", "-10",
		"120", "1201", "1202", "1203", "1204", "1205", "1206", "1207", "1208", "1209", "1210",
		"hello", "world", " ", "hello world hello", "hello world  ",
	}
	for _, s := range table {
		challenge.Program("getarea", strings.Fields(s)...)
	}
}
