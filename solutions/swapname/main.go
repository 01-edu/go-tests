package main

import (
	"fmt"
	"strings"

	"os"
)

func isValidName(fullName string) []string {
	result := make([]string, 0)
	lastpost := 0
	if len(fullName) < 2 {
		return nil
	}
	for i := 0; i < len(fullName); i++ {
		if fullName[i] >= 'a' && fullName[i] <= 'z' || fullName[i] >= 'A' && fullName[i] <= 'Z' || fullName[i] <= 32 {
			if fullName[i] == ' ' {
				if fullName[lastpost:i] != "" {
					result = append(result, fullName[lastpost:i])
				}
				lastpost = i + 1
			}
		} else {
			return nil
		}
	}
	result = append(result, fullName[lastpost:])
	return result
}

func main() {
	argument := os.Args[1:]
	if len(argument) == 1 {
		fullName := isValidName(strings.TrimSpace(argument[0]))
		if fullName != nil && len(fullName) == 2 {
			fmt.Println(fullName[1] + " " + fullName[0])
		} else {
			fmt.Println("Error")
		}
	}
}
