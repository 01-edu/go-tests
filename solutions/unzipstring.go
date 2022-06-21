package solutions

import "strconv"

func Unzipstring(str string) string {
    var result string
	if len(str) == 0 ||  len(str) < 2{
		return "Invalid output"
	}
    num, err := strconv.Atoi(string(str[0]))
    if err != nil {
        return "Invalid output"
    }
    for i := 0; i < num; i++ {
        result += string(str[1])
    }
    if len(str) > 2 {
        if Unzipstring(string(str[2:])) != "Invalid output" {
            result += Unzipstring(string(str[2:]))
        } else {
            return "Invalid output"
        }
    }
    return result
}
