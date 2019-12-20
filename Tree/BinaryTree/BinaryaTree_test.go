package BinaryTree

import (
	"fmt"
	"testing"
)

func GetTestSTree() *TreeNode {
	//st := STree{}
	/*
		   a
		  /  \
		 b    c
		  \   /
		  e   d
		 / \
		f  g
	*/
	// in b f e g a d c
	a := &TreeNode{Val: 1, Left: nil, Right: nil}
	b := &TreeNode{Val: 2, Left: nil, Right: nil}
	c := &TreeNode{Val: 3, Left: nil, Right: nil}
	d := &TreeNode{Val: 4, Left: nil, Right: nil}
	e := &TreeNode{Val: 5, Left: nil, Right: nil}
	f := &TreeNode{Val: 6, Left: nil, Right: nil}
	g := &TreeNode{Val: 7, Left: nil, Right: nil}

	a.Left = b
	a.Right = c
	c.Left = d
	b.Right = e
	e.Left = f
	e.Right = g

	return a
}

func TestSTree_ShowTree(t *testing.T) {
	st := STree{}
	testTree := GetTestSTree()
	st.ShowTree(testTree)
}

func TestMaxDepth(t *testing.T) {
	st := STree{}
	testTree := GetTestSTree()
	depth := st.MaxDepth(testTree)
	fmt.Print("depth: ", depth)
}

func TestF(t *testing.T) {
	L := []int{1, 2, 3, 4, 5}
	l := L[1:1]
	l1 := L[2:]
	fmt.Print(l, "\n", l1)
	//x := []int{}
	//x := make([]int,0)
	var x []int
	if x == nil {
		fmt.Print("x==nil")
	}
	if len(x) == 0 {
		fmt.Print("len(x)==0")
	}

}

func TestSTree_BuildTree(t *testing.T) {
	st := STree{}
	//testTree := GetTestSTree()
	preorder := []int{3, 9, 20, 15, 7}
	inorder := []int{9, 3, 15, 20, 7}
	st.BuildTree(preorder, inorder)
}
