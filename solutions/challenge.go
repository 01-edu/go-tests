package solutions

import (
	"strconv"

	"github.com/01-edu/go-tests/lib/challenge"
	"github.com/01-edu/go-tests/lib/chars"
	"github.com/01-edu/go-tests/lib/random"
)

func FormatTree(root *TreeNode) string {
	if root == nil {
		return ""
	}
	res := root.Data + "\n"
	res += formatSubTree(root, "")
	return res
}

func formatSubTree(root *TreeNode, prefix string) string {
	if root == nil {
		return ""
	}

	var res string

	hasLeft := root.Left != nil
	hasRight := root.Right != nil

	if !hasLeft && !hasRight {
		return res
	}

	res += prefix
	if hasLeft && hasRight {
		res += "├── "
	}

	if !hasLeft && hasRight {
		res += "└── "
	}

	if hasRight {
		printStrand := (hasLeft && hasRight && (root.Right.Right != nil || root.Right.Left != nil))
		newPrefix := prefix
		if printStrand {
			newPrefix += "│   "
		} else {
			newPrefix += "    "
		}
		res += root.Right.Data + "\n"
		res += formatSubTree(root.Right, newPrefix)
	}

	if hasLeft {
		if hasRight {
			res += prefix
		}
		res += "└── " + root.Left.Data + "\n"
		res += formatSubTree(root.Left, prefix+"    ")
	}
	return res
}

func ParentList(root *TreeNode) string {
	if root == nil {
		return ""
	}

	var parent string

	if root.Parent == nil {
		parent = "nil"
	} else {
		parent = root.Parent.Data
	}

	r := "Node: " + root.Data + " Parent: " + parent + "\n"
	r += ParentList(root.Left) + ParentList(root.Right)
	return r
}

func ChallengeTree(
	name string,
	fn1, fn2 interface{},
	arg1 *TreeNode, arg2 interface{},
	args ...interface{},
) {
	args1 := []interface{}{arg1}
	args2 := []interface{}{arg2}

	if args != nil {
		for _, v := range args {
			args1 = append(args1, v)
			args2 = append(args2, v)
		}
	}
	st1 := challenge.Monitor(fn1, args1)
	st2 := challenge.Monitor(fn2, args2)

	if st1.Stdout != st2.Stdout {
		challenge.Fatalf("%s(\n%s)\n prints %s instead of %s\n",
			name,
			FormatTree(arg1),
			challenge.Format(st2.Stdout),
			challenge.Format(st1.Stdout),
		)
	}
}

// tests the lists, by comparing values in both lists
func ChallengeList(f string, l, l1 *List, data []interface{}, args ...interface{}) {
	student := CopyList(l1)
	solution := CopyList(l)
	for l.Head != nil || l1.Head != nil {
		if l.Head == nil || l1.Head == nil {
			challenge.Fatalf("\ndata used: %v\nstudent list:%s\nlist:%s\n\n%s(%s)== %v instead of %v\n\n",
				data,
				ListToString(student.Head),
				ListToString(solution.Head),
				f,
				challenge.Format(args...),
				l1.Head,
				l.Head)
		}
		if l.Head.Data != l1.Head.Data {
			challenge.Fatalf("\ndata used: %v\nstudent list:%s\nlist:%s\n\n%s(%s)== %v instead of %v\n\n",
				data,
				ListToString(student.Head),
				ListToString(solution.Head),
				f,
				challenge.Format(args...),
				l1.Head.Data,
				l.Head.Data)
		}
		l.Head = l.Head.Next
		l1.Head = l1.Head.Next
	}
}

func CopyList(listStu *List) *List {
	copy := &List{}
	it := listStu.Head
	for it != nil {
		ListPushBack(copy, it.Data)
		it = it.Next
	}
	return copy
}

func ListToString(n *NodeL) string {
	var res string
	it := n
	for it != nil {
		switch it.Data.(type) {
		case int:
			res += strconv.Itoa(it.Data.(int)) + "-> "
		case string:
			res += it.Data.(string) + "-> "
		}
		it = it.Next
	}
	res += "<nil>"
	return res
}

func ListPushNode(l *NodeI, data int) *NodeI {
	n := &NodeI{Data: data}

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

func CopyNode(listStu *NodeI) *NodeI {
	var listSol *NodeI
	it := listStu
	for it != nil {
		listSol = ListPushNode(listSol, it.Data)
		it = it.Next
	}
	return listSol
}

func NodeToString(n *NodeI) string {
	var res string
	it := n
	for it != nil {
		res += strconv.Itoa(it.Data) + "-> "
		it = it.Next
	}
	res += "<nil>"
	return res
}

func IntToInterface(t []int) []interface{} {
	RandLen := random.IntBetween(0, len(t))
	s := make([]interface{}, RandLen)
	for j := 0; j < RandLen; j++ {
		for i := 0; i < random.IntBetween(1, len(t)); i++ {
			s[j] = t[i]
		}
	}
	return s
}

func IntToStringface(t []string) []interface{} {
	RandLen := random.IntBetween(0, len(t))
	s := make([]interface{}, RandLen)
	for j := 0; j < RandLen; j++ {
		for i := 0; i < random.IntBetween(1, len(t)); i++ {
			s[j] = t[i]
		}
	}
	return s
}

type NodeTest struct {
	Data []interface{}
}

func ElementsToTest(table []NodeTest) []NodeTest {
	table = append(table,
		NodeTest{
			Data: []interface{}{},
		},
	)
	for i := 0; i < 3; i++ {
		val := NodeTest{
			Data: IntToInterface(random.IntSliceBetween(-50, 100)),
		}
		table = append(table, val)
	}
	for i := 0; i < 3; i++ {
		val := NodeTest{
			Data: IntToStringface(random.StrSlice(chars.Words)),
		}
		table = append(table, val)
	}
	return table
}
