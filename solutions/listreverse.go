package solutions

func ListReverse(l *List) {
	current := l.Head
	var prev *NodeL
	prev = nil

	for current != nil {
		next := current.Next
		current.Next = prev
		prev = current
		current = next
	}
	l.Head = prev
}
