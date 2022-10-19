package solutions

func StringToBool(s string) bool {
	if s == "True" || s == "T" || s == "t" {
		return true
	}
	return false
}
