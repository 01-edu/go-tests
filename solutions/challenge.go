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
func ChallengeList(f string, a, b *List, data []interface{}, args ...interface{}) {
	student := copyList(b)
	solution := copyList(a)
	for a.Head != nil || b.Head != nil {
		if a.Head == nil || b.Head == nil {
			challenge.Fatalf("\ndata used: %v\nstudent list:%s\nlist:%s\n\n%s(%s)== %v instead of %v\n\n",
				data,
				ListToString(student.Head),
				ListToString(solution.Head),
				f,
				challenge.Format(args...),
				b.Head,
				a.Head)
		}
		if a.Head.Data != b.Head.Data {
			challenge.Fatalf("\ndata used: %v\nstudent list:%s\nlist:%s\n\n%s(%s)== %v instead of %v\n\n",
				data,
				ListToString(student.Head),
				ListToString(solution.Head),
				f,
				challenge.Format(args...),
				b.Head.Data,
				a.Head.Data)
		}
		a.Head = a.Head.Next
		b.Head = b.Head.Next
	}
}

func copyList(listStu *List) *List {
	copy := &List{}
	for it := listStu.Head; it != nil; it = it.Next {
		ListPushBack(copy, it.Data)
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

func ListPushNode(l *NodeI, data int) {
	n := &NodeI{Data: data}

	iterator := l
	if iterator == nil {
		iterator = n
		return
	}
	for iterator.Next != nil {
		iterator = iterator.Next
	}
	iterator.Next = n
}

func CopyNode(listStu *NodeI) (listSol *NodeI) {
	for it := listStu; it != nil; it = it.Next {
		ListPushNode(listSol, it.Data)
	}
	return
}

func NodeToString(n *NodeI) string {
	var res string
	for it := n; it != nil; it = it.Next {
		res += strconv.Itoa(it.Data) + "-> "
	}
	res += "<nil>"
	return res
}

func IntToInterface(slice []int) (a []interface{}) {
	for _, v := range slice {
		a = append(a, v)
	}
	return
}

func StringToInterface(slice []string) (a []interface{}) {
	for _, v := range slice {
		a = append(a, v)
	}
	return
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
		table = append(table, NodeTest{
			Data: IntToInterface(random.IntSliceBetween(-50, 100)),
		})
	}
	for i := 0; i < 3; i++ {
		table = append(table, NodeTest{
			Data: StringToInterface(random.StrSlice(chars.Words)),
		})
	}

	return table
}
