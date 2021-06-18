package main

import (
	"math/rand"
	"strconv"
	student "student"

	"github.com/01-edu/go-tests/lib/challenge"
	"github.com/01-edu/go-tests/lib/random"

	"github.com/01-edu/go-tests/solutions"
)

type (
	stuTreeNode = student.TreeNodeL
	solTreeNode = solutions.TreeNodeL
)

func insertStu(t *stuTreeNode, v int) *stuTreeNode {
	if t == nil {
		return &stuTreeNode{Left: nil, Val: v, Right: nil}
	}
	if v < t.Val {
		t.Left = insertStu(t.Left, v)
		return t
	}
	t.Right = insertStu(t.Right, v)
	return t
}

func compare(stuResult *stuTreeNode, solResult *solTreeNode) bool {
	c1, c2 := stuWalker(stuResult), solWalker(solResult)
	for {
		v1, ok1 := <-c1
		v2, ok2 := <-c2
		if !ok1 || !ok2 {
			return ok1 == ok2
		}
		if v1 != v2 {
			break
		}
	}
	return false
}

// Walk traverses a tree depth-first,
// sending each Value on a channel.
func stuWalk(t *stuTreeNode, ch chan int) {
	if t == nil {
		return
	}
	stuWalk(t.Left, ch)
	ch <- t.Val
	stuWalk(t.Right, ch)
}

// Walk traverses a tree depth-first,
// sending each Val on a channel.
func solWalk(t *solTreeNode, ch chan int) {
	if t == nil {
		return
	}
	solWalk(t.Left, ch)
	ch <- t.Val
	solWalk(t.Right, ch)
}

// Walker launches Walk in a new goroutine,
// and returns a read-only channel of values.
func stuWalker(t *stuTreeNode) <-chan int {
	ch := make(chan int)
	go func() {
		stuWalk(t, ch)
		close(ch)
	}()
	return ch
}

// Walker launches Walk in a new goroutine,
// and returns a read-only channel of values.
func solWalker(t *solTreeNode) <-chan int {
	ch := make(chan int)
	go func() {
		solWalk(t, ch)
		close(ch)
	}()
	return ch
}
func insertSol(t *solTreeNode, v int) *solTreeNode {
	if t == nil {
		return &solTreeNode{Left: nil, Val: v, Right: nil}
	}
	if v < t.Val {
		t.Left = insertSol(t.Left, v)
		return t
	}
	t.Right = insertSol(t.Right, v)
	return t
}

func New(n, k int) (*solTreeNode, *stuTreeNode, *solTreeNode) {
	var cop *solTreeNode
	var stu *stuTreeNode
	var sol *solTreeNode
	for _, v := range rand.Perm(n) {
		cop = insertSol(cop, (1+v)*k)
		stu = insertStu(stu, (1+v)*k)
		sol = insertSol(sol, (1+v)*k)
	}
	return cop, stu, sol
}

func returnSolTree(root *solTreeNode) string {
	if root == nil {
		return ""
	}
	ans := strconv.Itoa(root.Val)
	if root.Left == nil && root.Right == nil {
		return ans
	}
	if root.Left != nil {
		ans += " " + returnSolTree(root.Left)
	}
	if root.Right != nil {
		ans += " " + returnSolTree(root.Right)
	}
	return ans
}

func returnStuTree(root *stuTreeNode) string {
	if root == nil {
		return ""
	}
	ans := strconv.Itoa(root.Val)
	if root.Left == nil && root.Right == nil {
		return ans
	}
	if root.Left != nil {
		ans += " " + returnStuTree(root.Left)
	}
	if root.Right != nil {
		ans += " " + returnStuTree(root.Right)
	}
	return ans
}

func compareResults(stuResult *stuTreeNode, solResult, solTree1, solTree2 *solTreeNode) {
	if !compare(stuResult, solResult) {
		tree1 := returnSolTree(solTree1)
		tree2 := returnSolTree(solTree2)
		stuTree := returnStuTree(stuResult)
		solTree := returnSolTree(solResult)
		challenge.Fatalf("\nIs same trees(\"%v\", \"%v\") == \"%v\" instead of \"%v\"\n\n",
			tree1, tree2, stuTree, solTree)
	}
}

func main() {
	type node struct {
		n int
		k int
	}

	table := []node{}
	for i := 0; i < 15; i++ {
		value := node{random.IntBetween(10, 15), random.IntBetween(1, 10)}
		table = append(table, value)
	}
	for _, arg := range table {
		_, stuTree1, solTree1 := New(arg.n, arg.k)
		_, stuTree2, solTree2 := New(arg.n, arg.k)
		stuResult := student.IsSameTree(stuTree1, stuTree2)
		solResult := solutions.IsSameTree(solTree1, solTree2)
		if solResult != stuResult {
			challenge.Fatalf("\"%v\" instead of \"%v\"\n\n", stuResult, solResult)
		}
	}
}
