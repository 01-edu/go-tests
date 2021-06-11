package main

import (
	student "student"

	"github.com/01-edu/go-tests/lib/challenge"
	"github.com/01-edu/go-tests/lib/random"
	"github.com/01-edu/go-tests/solutions"
)

func challengeNodeI(l1 *student.NodeI, l2 *solutions.NodeI, data [][]int) {
	student := copyList(l1)
	solution := solutions.CopyNode(l2)
	for l1 != nil || l2 != nil {
		if (l1 == nil && l2 != nil) || (l1 != nil && l2 == nil) {
			challenge.Fatalf("\ndata set:%d\nstudent list:%s\nlist:%s\n\nSortListInsert() == %v instead of %v\n\n",
				data, solutions.NodeToString(student), solutions.NodeToString(solution), l1, l2)
		}
		if l1.Data != l2.Data {
			challenge.Fatalf("\ndata set:%d\nstudent list:%s\nlist:%s\n\nSortListInsert() == %v instead of %v\n\n",
				data, solutions.NodeToString(student), solutions.NodeToString(solution), l1.Data, l2.Data)
		}
		l1 = l1.Next
		l2 = l2.Next
	}
}

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

func copyList(listStu *student.NodeI) *solutions.NodeI {
	var listSol *solutions.NodeI
	it := listStu
	for it != nil {
		listSol = solutions.ListPushNode(listSol, it.Data)
		it = it.Next
	}
	return listSol
}

func move(l *student.NodeI) *student.NodeI {
	p := l
	n := l.Next
	ret := n

	for n != nil && l.Data > n.Data {
		p = n
		n = n.Next
	}
	p.Next = l
	l.Next = n
	return ret
}

func listSort(l *student.NodeI) *student.NodeI {
	head := l
	if head == nil {
		return nil
	}
	head.Next = listSort(head.Next)

	if head.Next != nil && head.Data > head.Next.Data {
		head = move(head)
	}
	return head
}

func main() {
	var link1 *solutions.NodeI
	var link2 *solutions.NodeI
	var linkTest1 *student.NodeI
	var linkTest2 *student.NodeI

	type nodeTest struct {
		data1 []int
		data2 []int
	}

	table := []nodeTest{{
		data1: []int{},
		data2: []int{},
	}, {
		data1: []int{3, 5, 7},
		data2: []int{1, -2, 4, 6},
	}}

	for i := 0; i < 3; i++ {
		table = append(table, nodeTest{
			data1: random.IntSliceBetween(-50, 100),
			data2: random.IntSliceBetween(-50, 100),
		})
	}

	for _, arg := range table {
		for _, item := range arg.data1 {
			link1 = solutions.ListPushNode(link1, item)
			linkTest1 = listPushBack(linkTest1, item)
		}

		for _, item := range arg.data2 {
			link2 = solutions.ListPushNode(link2, item)
			linkTest2 = listPushBack(linkTest2, item)
		}

		linkTest1 = listSort(linkTest1)
		linkTest2 = listSort(linkTest2)
		link1 = solutions.ListSort(link1)
		link2 = solutions.ListSort(link2)

		solutionList := solutions.SortedListMerge(link1, link2)
		studentList := student.SortedListMerge(linkTest1, linkTest2)

		challengeNodeI(studentList, solutionList, [][]int{arg.data1, arg.data2})

		link1 = nil
		link2 = nil
		linkTest1 = nil
		linkTest2 = nil
	}
}
