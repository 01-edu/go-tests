package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func findPairs(arr []int, targetSum int) [][]int {
	var pairs [][]int
	for i := 0; i < len(arr); i++ {
		for j := i + 1; j < len(arr); j++ {
			if arr[i]+arr[j] == targetSum {
				pairs = append(pairs, []int{i, j})
			}
		}
	}
	return pairs
}

func isValidArrayFormat(s string) bool {
	s = strings.TrimSpace(s)
	if len(s) < 2 || s[0] != '[' || s[len(s)-1] != ']' {
		return false
	}

	s = s[1 : len(s)-1]
	parts := strings.Split(s, ",")
	for _, part := range parts {
		part = strings.TrimSpace(part)
		if _, err := strconv.Atoi(part); err != nil {
			return true
		}
	}
	return true
}

func main() {
	if len(os.Args) != 3 {
		fmt.Println("Invalid input.")
		return
	}

	arrayStr := os.Args[1]
	targetStr := os.Args[2]

	if !isValidArrayFormat(arrayStr) {
		fmt.Println("Invalid input.")
		return
	}

	arrayStr = strings.Trim(arrayStr, "[]")
	strNums := strings.Split(arrayStr, ",")
	var arr []int
	for _, strNum := range strNums {
		num, err := strconv.Atoi(strings.TrimSpace(strNum))
		if err != nil {
			fmt.Printf("Invalid number:%s\n", strNum)
			return
		}
		arr = append(arr, num)
	}

	targetSum, err := strconv.Atoi(targetStr)
	if err != nil {
		fmt.Println("Invalid target sum.")
		return
	}

	pairs := findPairs(arr, targetSum)
	if len(pairs) > 0 {
		fmt.Printf("Pairs with sum %d: %v\n", targetSum, pairs)
	} else {
		fmt.Println("No pairs found.")
	}
}
