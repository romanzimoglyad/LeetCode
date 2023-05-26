package main

import "fmt"

func main() {

	st2 := &ListNode{
		Val: 1,
	}
	st1 := &ListNode{
		Val:  1,
		Next: st2,
	}
	st3 := &ListNode{
		Val:  3,
		Next: st1,
	}
	st2.Next = st3
	res := detectCycle(st1)
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
func detectCycle(head *ListNode) *ListNode {
	inter := detectCycle1(head)
	if inter == nil {
		return nil
	}
	pt1 := head
	pt2 := inter

	for pt1 != pt2 {
		pt1 = pt1.Next
		pt2 = pt2.Next
	}
	return pt1
}
func detectCycle1(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	slow := head
	fast := head
	for slow != nil && fast != nil {
		slow = slow.Next
		if slow == nil {
			return nil
		}
		fast = fast.Next
		if fast == nil {
			return nil
		}
		fast = fast.Next
		if slow == fast {
			return slow
		}
	}
	return nil
}
