package solutions

import "strconv"

func countDuplication(s string, i int) int {
	var count int = 0
	for _, v := range s[i:] {
		if v == rune(s[i]) {
			count++
		} else {
			break
		}
	}
	return count
}

func ZipString(s string) string {
	var result string
	i := 0
	for i < len(s) {
		counter := countDuplication(s, i)
		result = result + strconv.Itoa(counter) + string(s[i])
		i += int(counter)
	}
	return result
}
