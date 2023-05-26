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
		Val:  1,
		Next: st1,
	}
	st2.Next = st3
	res := hasCycle(st1)
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
func hasCycle(head *ListNode) bool {
	if head == nil {
		return false
	}
	slow := head
	fast := head
	for slow != nil && fast != nil {
		slow = slow.Next
		if slow == nil {
			return false
		}
		fast = fast.Next
		if fast == nil {
			return false
		}
		fast = fast.Next
		if slow == fast {
			return true
		}
	}
	return false
}
