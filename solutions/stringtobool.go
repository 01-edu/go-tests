package solutions

func StringToBool(s string) bool {
	if s == "true" || s == "TRUE" || s == "T" || s == "t" {
		return true
	}
	return false
}
