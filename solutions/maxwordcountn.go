package solutions

import (
	"sort"
	"strings"
)

type keyValue struct {
	Key   string
	Value int
}

func MaxWordCountN(text string, n int) map[string]int {
	wordOccurences := make(map[string]int)
	for _, word := range strings.Split(text, " ") {
		wordOccurences[word]++
	}

	var keyValSlice []keyValue
	for k, v := range wordOccurences {
		keyValSlice = append(keyValSlice, keyValue{k, v})
	}

	sort.Slice(keyValSlice, func(i, j int) bool {
		if keyValSlice[i].Value == keyValSlice[j].Value {
			return keyValSlice[i].Key < keyValSlice[j].Key
		} else {
			return keyValSlice[i].Value > keyValSlice[j].Value
		}
	})

	nMaxWordOccurences := make(map[string]int)
	for _, kv := range keyValSlice[0:n] {
		nMaxWordOccurences[kv.Key] = kv.Value
	}
	return nMaxWordOccurences
}
