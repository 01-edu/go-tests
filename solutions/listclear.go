package solutions

func ListClear(l *List) {
	temp := l.Head
	var next *NodeL
	for temp != nil {
		next = temp.Next
		temp = nil
		temp = next
	}

	l.Head = nil
}
