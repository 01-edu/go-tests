package main

import (
	"fmt"

	"os"
)

func isValidName(fullName string) string {
	result := ""
	if len(fullName) < 2 {
		return ""
	}
	for i := 0; i < len(fullName); i++ {
		if fullName[i] >= 'a' && fullName[i] <= 'z' || fullName[i] >= 'A' && fullName[i] <= 'Z' {
			result += string(fullName[i])
		} else {
			return ""
		}
	}

	return result
}

func deleteSpace(fullName string) []string {
	str := []string{}
	nextSpace := 0
	for i := 0; i < len(fullName); i++ {
		if fullName[i] == ' ' {
			str = append(str, fullName[nextSpace:i])
			nextSpace = i + 1
		}
	}
	str = append(str, fullName[nextSpace:])
	return str
}

func main() {
	argument := os.Args[1:]
	if len(argument) == 1 {
		fullname := deleteSpace(argument[0])
		if len(fullname) != 2 {
			fmt.Println("Error")
			return
		}
		name := isValidName(fullname[0])
		lastname := isValidName(fullname[1])
		if name == "" || lastname == "" {
			fmt.Println("Error")
			return
		}
		fmt.Println(lastname + " " + name)
	}
}
