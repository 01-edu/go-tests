package solutions

func ReverseStrings(arr []string) string {
	var reversed string
	for i := len(arr) - 1; i >= 0; i-- {
		for j := len(arr[i]) - 1; j >= 0; j-- {
			reversed += string(arr[i][j])
		}
		if i != 0 && len(reversed) != 0 {
			reversed += " "
		}
	}
	return reversed
}
