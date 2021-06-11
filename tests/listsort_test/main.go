package main

import (
	"strconv"
	student "student"

	"github.com/01-edu/go-tests/lib/challenge"
	"github.com/01-edu/go-tests/lib/random"
	"github.com/01-edu/go-tests/solutions"
)

func listPushBack(l *student.NodeI, data int) {
	n := &student.NodeI{Data: data}
	if l == nil {
		l = n
	}
	iterator := l
	for iterator.Next != nil {
		iterator = iterator.Next
	}
	iterator.Next = n
}

func nodeToString(n *student.NodeI) string {
	var res string
	it := n
	for it != nil {
		res += strconv.Itoa(it.Data) + "-> "
		it = it.Next
	}
	res += "<nil>"
	return res
}

func challengeNodeI(a *student.NodeI, b *solutions.NodeI) {
	for a != nil || b != nil {
		if (a == nil && b != nil) || (a != nil && b == nil) {
			challenge.Fatalf("\nstudent list:%s\nlist:%s\n\nListSort() == %v instead of %v\n\n",
				nodeToString(a), solutions.NodeToString(b), a, b)
		}
		if a.Data != b.Data {
			challenge.Fatalf("\nstudent list:%s\nlist:%s\n\nListSort() == %v instead of %v\n\n",
				nodeToString(a), solutions.NodeToString(b), a.Data, b.Data)
		}
		a = a.Next
		b = b.Next
	}
}

func main() {
	var link1 *solutions.NodeI
	var link2 *student.NodeI

	type nodeTest struct {
		data []int
	}
	table := []nodeTest{
		{data: []int{}},
		{data: []int{5, 4, 3, 2, 1}},
	}
	for i := 0; i < 2; i++ {
		table = append(table, nodeTest{random.IntSliceBetween(-50, 100)})
	}

	for _, arg := range table {
		for _, item := range arg.data {
			listPushBack(link2, item)
			solutions.ListPushNode(link1, item)
		}

		sortedStudent := student.ListSort(link2)
		sortedSolution := solutions.ListSort(link1)

		challengeNodeI(sortedStudent, sortedSolution)
		link1 = &solutions.NodeI{}
		link2 = &student.NodeI{}
	}
}
