package Lists

import (
	"fmt"
	"testing"
)

func TestSList(t *testing.T) {
	s := &SList{}
	L := s.GetSList(5)
	s.ShowSLists(L)
}

func TestReverse(t *testing.T) {
	s := &SList{}
	L := s.GetSList(2)
	//L := &ListNode{
	//	Val: 0,
	//	Next: nil,
	//}
	r := s.Reverse(L)
	s.ShowSLists(r)
}

func TestCheckRang(t *testing.T) {
	s := &SList{}
	//L := s.GetSList(3)
	//L := &ListNode{
	//	Val: 0,
	//	Next: nil,
	//}
	//L.Next = L
	//s.ShowSLists(L)
	a := &ListNode{Val: 0, Next: nil}
	b := &ListNode{Val: 1, Next: nil}
	c := &ListNode{Val: 2, Next: nil}
	d := &ListNode{Val: 3, Next: nil}
	e := &ListNode{Val: 4, Next: nil}
	f := &ListNode{Val: 5, Next: nil}

	a.Next = b
	b.Next = c
	c.Next = d
	d.Next = e
	e.Next = f
	f.Next = d
	b.Next = b
	r := s.hasCycle(a)
	//s.ShowSLists(r)
	fmt.Println(r)
}
