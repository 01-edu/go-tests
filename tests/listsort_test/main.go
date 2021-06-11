package main

import (
	student "student"

	"github.com/01-edu/go-tests/lib/challenge"
	"github.com/01-edu/go-tests/lib/random"
	"github.com/01-edu/go-tests/solutions"
)

func listPushBack(l *student.NodeI, data int) *student.NodeI {
	n := &student.NodeI{Data: data}
	if l == nil {
		return n
	} else {
		iterator := l
		for iterator.Next != nil {
			iterator = iterator.Next
		}
		iterator.Next = n
	}
	return l
}

func challengeNodeI(l1 *student.NodeI, l2 *solutions.NodeI) {
	for l1 != nil || l2 != nil {
		if (l1 == nil && l2 != nil) || (l1 != nil && l2 == nil) {
			challenge.Fatalf("\nstudent list:%s\nlist:%s\n\nListSort() == %v instead of %v\n\n",
				solutions.NodeToString(copyList(l1)), solutions.NodeToString(l2), l1, l2)
		}
		if l1.Data != l2.Data {
			challenge.Fatalf("\nstudent list:%s\nlist:%s\n\nListSort() == %v instead of %v\n\n",
				solutions.NodeToString(copyList(l1)), solutions.NodeToString(l2), l1.Data, l2.Data)
		}
		l1 = l1.Next
		l2 = l2.Next
	}
}

func copyList(listStu *student.NodeI) *solutions.NodeI {
	var listSol *solutions.NodeI
	it := listStu
	for it != nil {
		listSol = solutions.ListPushNode(listSol, it.Data)
		it = it.Next
	}
	return listSol
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
			link2 = listPushBack(link2, item)
			link1 = solutions.ListPushNode(link1, item)
		}

		sortedStudent := student.ListSort(link2)
		sortedSolution := solutions.ListSort(link1)

		challengeNodeI(sortedStudent, sortedSolution)
		link1 = &solutions.NodeI{}
		link2 = &student.NodeI{}
	}
}
