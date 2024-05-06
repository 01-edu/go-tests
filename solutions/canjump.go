package solutions

func CanJump(steps []uint) bool {
	if len(steps) == 0 {
		return false
	}

	maxReach := 0
	for i := 0; i < len(steps); i++ {
		if i > maxReach {
			return false
		}
		maxReach = max(maxReach, i+int(steps[i]))
		if maxReach == len(steps)-1 || len(steps) == 1 {
			return true
		}
	}
	return false
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
