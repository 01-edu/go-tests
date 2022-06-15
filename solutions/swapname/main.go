package main

import (
	"fmt"

	"os"
)

func isValidName(fullName string) bool {
	if len(fullName) < 2 {
		return false
	}
	for i := 0; i < len(fullName); i++ {
		if fullName[i] >= 'a' && fullName[i] <= 'z' || fullName[i] >= 'A' && fullName[i] <= 'Z' || fullName[i] <= 32 {
			continue
		} else {
			return false
		}
	}
	return true
}

func findSpace(name string) bool {
	for i := 0; i < len(name); i++ {
		if name[i] == ' ' {
			return true
		}
	}
	return false
}

func swapName(fullName string) {
	fistName := ""
	lastName := ""
	for i := 0; i < len(fullName); i++ {
		if fullName[i] != ' ' {
			lastName += string(fullName[i])
		} else {
			fistName += fullName[i+1:]
			if findSpace(fistName){
				fmt.Println("Error")
				return
			}
			break
		}
	}
	fmt.Println(fistName + " " + lastName)
}

func main() {
	argument := os.Args[1:]
	if len(argument) == 1 && len(argument[0]) > 1 {
		for i := 0; i < len(argument[0]); i++ {
			if argument[0][i] == ' ' {
				continue
			}
			if !isValidName(argument[0][i:]) {
				fmt.Println("Error")
				return
			}
			swapName(argument[0][i:])
			break
		}
	}
}
