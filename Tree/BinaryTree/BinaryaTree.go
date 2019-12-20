package BinaryTree

import (
	"fmt"
	"math"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type STree struct {
}

/*
递归遍历思路
遍历Tree =
	子问题A 遍历左节点
	子问题B 遍历右节点
利用 A B 解决总问题

*/

type ShowTree struct {
}

// root -> left -> right
/*
前序遍历的递推公式：
pre = print root + pre(root.left) + pre(root.right)

*/
func (st *ShowTree) PreOrder(root *TreeNode) {
	if root == nil {
		return
	}
	fmt.Printf("%d ", root.Val)
	st.PreOrder(root.Left)
	st.PreOrder(root.Right)
}

// left -> root -> right
/*
中序遍历的递推公式：
pre = pre(root.left) + print root + pre(root.right)

*/
func (st *ShowTree) InOrder(root *TreeNode) {
	if root == nil {
		return
	}
	st.InOrder(root.Left)
	fmt.Printf("%d ", root.Val)
	st.InOrder(root.Right)
}

// left -> right -> root
/*
后序遍历的递推公式：
pre = pre(root.left) + pre(root.right) + print root

*/
func (st *ShowTree) PostOrder(root *TreeNode) {
	if root == nil {
		return
	}
	st.PostOrder(root.Left)
	st.PostOrder(root.Right)
	fmt.Printf("%d ", root.Val)
}

func (s *STree) ShowTree(root *TreeNode) {
	st := &ShowTree{}
	st.InOrder(root)
}

/*
递归遍历思路
maxDepth =
	子问题A 左节点深度
	子问题B 右节点深度
利用 A B 解决总问题

	1. 每递归一层 深度加一
	2. 取左/右深度大值

递归公式
	树深度=max(左子树深度,右子树深度) + 1
终止条件
	叶子节点不存在 return 0

*/
func (s *STree) MaxDepth(root *TreeNode) (depth int) {
	if root == nil {
		return 0
	}
	leftDepth := s.MaxDepth(root.Left)
	rightDepth := s.MaxDepth(root.Right)
	maxDepth := math.Max(float64(leftDepth), float64(rightDepth))
	return int(maxDepth) + 1
}

/*
preorder = [3,9,20,15,7]
inorder = [9,3,15,20,7]

    3
   / \
  9  20
    /  \
   15   7

*/

/*
递归思路
BuildTree =
	子问题A 左节点遍历
	子问题B 右节点遍历
利用 A B 解决总问题

	1. 分别解决 左节点 preorder inorder 的值
		利用 preorder 第一个必为 root 再通过 inorder 获取 下一次递归 inorder 的左节点以及右节点，
		以及 preorder 的左右节点数量得到下一次递归 preorder

总结
	理解各个遍历的特点
*/

func (s *STree) BuildTree(preorder []int, inorder []int) *TreeNode {
	if (preorder == nil) || (inorder == nil) || len(preorder) == 0 || len(inorder) == 0 {
		return nil
	}
	fmt.Println("preorder[0]", preorder[0])
	root := &TreeNode{
		Val:   preorder[0],
		Left:  nil,
		Right: nil,
	}
	var inRootIndex = 0
	for i := 0; i < len(inorder); i++ {
		if inorder[i] == preorder[0] {
			inRootIndex = i
			break
		}
	}

	fmt.Println("inRootIndex: ", inRootIndex)
	root.Left = s.BuildTree(preorder[1:inRootIndex+1], inorder[:inRootIndex+1])
	root.Right = s.BuildTree(preorder[inRootIndex+1:], inorder[inRootIndex+1:])
	return root
}
