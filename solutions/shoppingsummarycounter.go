package solutions

import "strings"

func ShoppingSummaryCounter(str string) map[string]int {
	wordList := strings.Split(str, " ")
	counts := make(map[string]int)
	for _, word := range wordList {
		_, ok := counts[word]
		if ok {
			counts[word] += 1
		} else {
			counts[word] = 1
		}
	}
	return counts
}
