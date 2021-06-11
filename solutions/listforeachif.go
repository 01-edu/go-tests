package solutions

// compare each element of the linked list to see if it is a String
func IsPositiveNode(node *NodeL) bool {
	switch node.Data.(type) {
	case int, float32, float64, byte:
		return node.Data.(int) > 0
	default:
		return false
	}
}

func IsAlNode(node *NodeL) bool {
	switch node.Data.(type) {
	case int, float32, float64, byte:
		return false
	default:
		return true
	}
}

// applies the function f on each string if the boolean function comp returns true
func ListForEachIf(l *List, f func(*NodeL), comp func(*NodeL) bool) {
	it := l.Head
	for it != nil {
		if comp(it) {
			f(it)
		}
		it = it.Next
	}
}
