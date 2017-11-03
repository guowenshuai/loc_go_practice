package main

import (
	"fmt"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	fmt.Println("Hello, playground")
	l1 := &ListNode{5, &ListNode{4, &ListNode{4, nil}}}
	l2 := &ListNode{5, &ListNode{8, &ListNode{4, nil}}}
	r := addTwoNumbers(l1, l2)
	fmt.Println("R:", r.Val, r.Next.Val, r.Next.Next.Val)

}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {

	ret := &ListNode{0, nil}
	retNext := ret
	over := 0

ADD:
	sum := l1.Val + l2.Val + over
	retNext.Val = sum % 10
	over = sum / 10

	if l1.Next == nil && l2.Next == nil {
		if over > 0 {
			retNext.Next = &ListNode{over, nil}
		}
		return ret
	}
	//	for _, lv := range []*ListNode{l1, l2} {
	//		if lv.Next == nil {
	//			lv.Next = &ListNode{0, nil}
	//		}
	//		lv = lv.Next
	//	}
	if l1.Next == nil {
		l1.Next = &ListNode{0, nil}
	}
	l1 = l1.Next
	if l2.Next == nil {
		l2.Next = &ListNode{0, nil}
	}
	l2 = l2.Next

	retNext.Next = &ListNode{0, nil}
	retNext = retNext.Next
	goto ADD

	//	return ret
}
