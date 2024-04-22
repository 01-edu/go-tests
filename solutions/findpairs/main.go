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

func main() {
	if len(os.Args) != 3 {
		fmt.Println("Usage: go run . [array] <targetSum>")
		return
	}

	arrayStr := os.Args[1]
	arrayStr = strings.TrimPrefix(arrayStr, "[")
	arrayStr = strings.TrimSuffix(arrayStr, "]")
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

	targetSum, err := strconv.Atoi(os.Args[2])
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
