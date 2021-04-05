package solutions

import "github.com/01-edu/go-tests/lib/base"

func convertNbrBase(n int, base string) string {
	var result string
	length := len(base)

	for n >= length {
		result = string(base[(n%length)]) + result
		n /= length
	}
	result = string(base[n]) + result

	return result
}

func ConvertBase(nbr, baseFrom, baseTo string) string {
	resultIntermediary := base.Atoi(nbr, baseFrom)

	resultFinal := convertNbrBase(resultIntermediary, baseTo)

	return resultFinal
}
