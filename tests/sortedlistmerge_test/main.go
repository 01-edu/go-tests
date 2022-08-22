package main

import (
	"strconv"
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
	}}

	for i := 0; i < 3; i++ {
		table = append(table, nodeTest{
			data1: random.IntSliceBetween(-50, 100),
			data2: random.IntSliceBetween(-50, 100),
		})
	}

	for _, arg := range table {
		for _, item := range arg.data1 {
			solutions.ListPushNode(link1, item)
			linkTest1 = listPushBack(linkTest1, item)
		}

		for _, item := range arg.data2 {
			solutions.ListPushNode(link2, item)
			linkTest2 = listPushBack(linkTest2, item)
		}

		linkTest1 = student.ListSort(linkTest1)
		linkTest2 = student.ListSort(linkTest2)
		link1 = solutions.ListSort(link1)
		link2 = solutions.ListSort(link2)

		solutionList := solutions.SortedListMerge(link1, link2)
		studentList := student.SortedListMerge(linkTest1, linkTest2)

		data := [][]int{arg.data1, arg.data2}
		for studentList != nil || solutionList != nil {
			if (studentList == nil && solutionList != nil) || (studentList != nil && solutionList == nil) {
				challenge.Fatalf("\ndata set:%d\nstudent list:%s\nlist:%s\n\nSortListInsert() == %v instead of %v\n\n",
					data, nodeToString(studentList), solutions.NodeToString(solutionList), studentList, solutionList)
			}
			if studentList.Data != solutionList.Data {
				challenge.Fatalf("\ndata set:%d\nstudent list:%s\nlist:%s\n\nSortListInsert() == %v instead of %v\n\n",
					data, nodeToString(studentList), solutions.NodeToString(solutionList), studentList.Data, solutionList.Data)
			}
			studentList = studentList.Next
			solutionList = solutionList.Next
		}

		link1 = &solutions.NodeI{}
		link2 = &solutions.NodeI{}
		linkTest1 = &student.NodeI{}
		linkTest2 = &student.NodeI{}
	}
}
