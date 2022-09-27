package solutions

func PodiumPosition(podium [][]string) [][]string {
	for i, j := 0, len(podium) - 1; i < j; i, j = i + 1, j - 1 {
        podium[i], podium[j] = podium[j], podium[i]
    }
	return podium
}
