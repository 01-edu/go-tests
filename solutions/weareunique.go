package solutions

import "strings"

func WeAreUnique(str1 string, str2 string) int {
	var used [127]int
	if str1 == "" && str2 == "" {
		return -1
	}
	var argv []string = []string{str1, str2}
	k := 0
	i := 0
	for i < 2 {
		j := 0
		for j < len(argv[i]) {
			if used[argv[i][j]] == 0 && !strings.Contains(argv[1-i], string(argv[i][j])) {
				used[argv[i][j]] = 1
				k++
			}
			j++
		}
		i++
	}
	return k
}
