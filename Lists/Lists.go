package Lists

import "fmt"

//206，141，21，19，876

//单链表反转
//
//链表中环的检测
//
//两个有序的链表合并
//
//删除链表倒数第n个结点
//
//求链表的中间结点

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {

}

type SList struct {
}

func (s *SList) GetSList(num int) *ListNode {
	var pre, cur *ListNode
	Head := &ListNode{
		Val:  0,
		Next: nil,
	}
	pre = Head
	if num == 1 {
		return pre
	}
	for i := 1; i < num; i++ {
		node := &ListNode{
			Val:  i,
			Next: nil,
		}
		//fmt.Println(node.Val)
		cur = node
		//Next = node
		pre.Next = cur
		pre = cur
		//cur = Next
	}
	return Head
}

func (s *SList) ShowSLists(head *ListNode) {
	var node *ListNode
	node = head
	fmt.Println("Head: ", node.Val)
	for node.Next != nil {
		fmt.Println(node.Val)
		node = node.Next
	}
	if node.Next == nil {
		fmt.Println(node.Val)
		return
	}
}

// 0>1>2>3>4>5
// 0<1<2<3<4<5

func (s *SList) Reverse(head *ListNode) *ListNode {
	var pre, cur, next *ListNode
	//创建哨兵节点
	ListNode := &ListNode{
		Val:  999,
		Next: nil,
	}
	//边界检查
	if head == nil {
		return head
	}
	if head.Next == nil {
		return head
	}
	// 创建三个指针，方便移动并记录
	pre = ListNode
	cur = head
	next = cur.Next

	//边界检查
	if cur.Next.Next == nil {
		next.Next = head
		cur.Next = nil
		return next
	}
	//循环
	for next.Next != nil {
		//翻转
		cur.Next = pre
		//三头移动
		pre = cur
		cur = next
		next = next.Next
	}

	//边界检查 最后操作和之前不一样
	if next.Next == nil {
		//翻转
		cur.Next = pre
		//最后节点再翻转
		next.Next = cur
	}
	//去掉哨兵节点
	head.Next = nil
	return next
}

// 1>2>3>4>5
//      /  \
//       7<6

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func (s *SList) hasCycle(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return false
	}
	if head.Next == head {
		return true
	}
	if head.Next.Next == head {
		return true
	}
	if head.Next.Next == nil {
		return false
	}
	var one, two *ListNode
	one, two = head, head
	for two.Next.Next != nil {
		one = one.Next
		two = two.Next.Next
		if one == two {
			return true
		}
		if two.Next == nil {
			return false
		}
	}
	return false
}
