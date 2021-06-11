package main

import (
	student "student"

	"github.com/01-edu/go-tests/lib/challenge"
	"github.com/01-edu/go-tests/lib/random"
	"github.com/01-edu/go-tests/solutions"
)

func challengeNodeI(a *student.NodeI, b, c *solutions.NodeI, data []int) {
	solution := solutions.CopyNode(b)
	student := copyList(a)
	for a != nil || b != nil {
		if (a == nil && b != nil) || (a != nil && b == nil) {
			challenge.Fatalf("\nlist before insert:%s\ndata used to insert: %d\nstudent list:%s\nlist:%s\n\nSortListInsert() == %v instead of %v\n\n",
				solutions.NodeToString(c), data, solutions.NodeToString(student), solutions.NodeToString(solution), a, b)
		}
		if a.Data != b.Data {
			challenge.Fatalf("\nlist before insert:%s\ndata used to insert: %d\nstudent list:%s\nlist:%s\n\nSortListInsert() == %v instead of %v\n\n",
				solutions.NodeToString(c), data, solutions.NodeToString(student), solutions.NodeToString(solution), a.Data, b.Data)
		}
		a = a.Next
		b = b.Next
	}
}

func copyList(listStu *student.NodeI) (listSol *solutions.NodeI) {
	for it := listStu; it != nil; it = it.Next {
		solutions.ListPushNode(listSol, it.Data)
	}
	return
}

func listPushBack(l *student.NodeI, data int) {
	n := &student.NodeI{Data: data}
	iterator := l
	for iterator.Next != nil {
		iterator = iterator.Next
	}
	iterator.Next = n
	l = iterator
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
	var view *solutions.NodeI
	var link2 *student.NodeI

	type nodeTest struct {
		data     []int
		data_ref []int
	}
	table := []nodeTest{{
		data:     []int{},
		data_ref: []int{},
	}}
	table = append(table,
		nodeTest{
			data:     []int{5, 4, 3, 2, 1},
			data_ref: random.IntSliceBetween(-50, 100),
		},
	)

	for i := 0; i < 2; i++ {
		table = append(table, nodeTest{
			data:     random.IntSliceBetween(-50, 100),
			data_ref: random.IntSliceBetween(-50, 100),
		})
	}
	for _, arg := range table {
		for _, item := range arg.data {
			solutions.ListPushNode(link1, item)
			listPushBack(link2, item)
			solutions.ListPushNode(view, item)
		}
		link1 = solutions.ListSort(link1)
		view = solutions.ListSort(view)
		link2 = listSort(link2)

		for _, item := range arg.data_ref {
			link1 = solutions.SortListInsert(link1, item)
			link2 = student.SortListInsert(link2, item)
		}

		challengeNodeI(link2, link1, view, arg.data_ref)
		link1 = &solutions.NodeI{}
		view = &solutions.NodeI{}
		link2 = &student.NodeI{}
	}
}
