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
	/*	st3 := &ListNode{
			Val: 3,
		}
		st2.Next = st3*/

	/*	st5 := &ListNode{
			Val: 1,
		}
		st4 := &ListNode{
			Val:  4,
			Next: st5,
		}

		st5.Next = st2*/

	res := removeNthFromEnd(st1, 2)
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
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	cur1 := head
	i := 0
	cur2 := head
	for i < n {
		cur1 = cur1.Next
		i++
	}
	for cur1 != nil && cur1.Next != nil {
		cur1 = cur1.Next
		cur2 = cur2.Next
	}
	if cur1 == nil {
		return head.Next
	}
	cur2.Next = cur2.Next.Next
	return head
}
