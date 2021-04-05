package main

import (
	"strconv"

	student "student"

	"github.com/01-edu/go-tests/lib/challenge"
	"github.com/01-edu/go-tests/lib/chars"
	"github.com/01-edu/go-tests/lib/random"
	"github.com/01-edu/go-tests/solutions"
)

type (
	Node11  = student.NodeL
	List11  = solutions.List
	NodeS11 = solutions.NodeL
	ListS11 = student.List
)

func listToStringStu(l *ListS11) string {
	var res string
	it := l.Head
	for it != nil {
		switch it.Data.(type) {
		case int:
			res += strconv.Itoa(it.Data.(int)) + "->"
		case string:
			res += it.Data.(string) + "->"
		}
		it = it.Next
	}
	res += "<nil>"
	return res
}

func listPushBackTest11(l1 *ListS11, l2 *List11, data interface{}) {
	n1 := &Node11{Data: data}
	n2 := &NodeS11{Data: data}

	if l1.Head == nil {
		l1.Head = n1
	} else {
		iterator := l1.Head
		for iterator.Next != nil {
			iterator = iterator.Next
		}
		iterator.Next = n1
	}
	l1.Tail = n1

	if l2.Head == nil {
		l2.Head = n2
	} else {
		iterator1 := l2.Head
		for iterator1.Next != nil {
			iterator1 = iterator1.Next
		}
		iterator1.Next = n2
	}
	l2.Tail = n2
}

func comparFuncList11(l1 *List11, l2 *ListS11) {
	for l1.Head != nil || l2.Head != nil {
		if (l1.Head == nil && l2.Head != nil) || (l1.Head != nil && l2.Head == nil) {
			challenge.Fatalf("\nstudent list:%s\nlist:%s\n\nListMerge() == %v instead of %v\n\n",
				listToStringStu(l2), solutions.ListToString(l1.Head), l2.Head, l1.Head)
		}
		if l1.Head.Data != l2.Head.Data {
			challenge.Fatalf("\nstudent list:%s\nlist:%s\n\nListMerge() == %v instead of %v\n\n",
				listToStringStu(l2), solutions.ListToString(l1.Head), l2.Head.Data, l1.Head.Data)
		}
		l1.Head = l1.Head.Next
		l2.Head = l2.Head.Next
	}
}

func main() {
	link1 := &List11{}
	linkTest := &List11{}
	link2 := &ListS11{}
	link2Test := &ListS11{}

	type nodeTest struct {
		data1 []interface{}
		data2 []interface{}
	}
	table := []nodeTest{}
	// empty list
	table = append(table,
		nodeTest{
			data1: []interface{}{},
			data2: []interface{}{},
		})
	table = append(table,
		nodeTest{
			data1: solutions.ConvertIntToInterface(random.IntSlice()),
			data2: []interface{}{},
		})
	// jut ints
	for i := 0; i < 3; i++ {
		val := nodeTest{
			data1: solutions.ConvertIntToInterface(random.IntSlice()),
			data2: solutions.ConvertIntToInterface(random.IntSlice()),
		}
		table = append(table, val)
	}
	// just strings
	for i := 0; i < 2; i++ {
		val := nodeTest{
			data1: solutions.ConvertIntToStringface(random.StrSlice(chars.Words)),
			data2: solutions.ConvertIntToStringface(random.StrSlice(chars.Words)),
		}
		table = append(table, val)
	}
	table = append(table,
		nodeTest{
			data1: []interface{}{},
			data2: []interface{}{"a", 1},
		},
	)
	for _, arg := range table {
		for i := 0; i < len(arg.data1); i++ {
			listPushBackTest11(link2, link1, arg.data1[i])
		}
		for i := 0; i < len(arg.data2); i++ {
			listPushBackTest11(link2Test, linkTest, arg.data2[i])
		}

		solutions.ListMerge(link1, linkTest)
		student.ListMerge(link2, link2Test)
		comparFuncList11(link1, link2)

		link1 = &List11{}
		linkTest = &List11{}
		link2 = &ListS11{}
		link2Test = &ListS11{}
	}
}
