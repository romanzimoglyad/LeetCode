package main

import "fmt"

func main() {

	st2 := &ListNode{
		Val: 2,
	}
	st1 := &ListNode{
		Val:  1,
		Next: st2,
	}
	st3 := &ListNode{
		Val: 3,
	}
	st2.Next = st3

	st5 := &ListNode{
		Val: 5,
	}
	st4 := &ListNode{
		Val:  4,
		Next: st5,
	}

	res := getIntersectionNode(st1, st4)
	fmt.Println(res)
}

type ListNode struct {
	Val  int
	Next *ListNode
}

/**
 * Definition for singly-linked list.
 *
 */
func getIntersectionNode(headA, headB *ListNode) *ListNode {
	cur1 := headA
	cur2 := headB

	for cur1 != cur2 {
		if cur1 == nil {
			cur1 = headB
		} else {
			cur1 = cur1.Next
		}
		if cur2 == nil {
			cur2 = headA
		} else {
			cur2 = cur2.Next
		}
	}
	return cur1
}
