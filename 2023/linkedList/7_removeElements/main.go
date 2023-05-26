package main

import "fmt"

func main() {

	st2 := &ListNode{
		Val: 7,
	}
	st1 := &ListNode{
		Val:  7,
		Next: st2,
	}
	st3 := &ListNode{
		Val: 7,
	}
	st2.Next = st3

	/*	st5 := &ListNode{
			Val: 1,
		}
		st4 := &ListNode{
			Val:  4,
			Next: st5,
		}

		st5.Next = st2*/

	res := removeElements(st1, 7)
	fmt.Println(res)
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func removeElements(head *ListNode, val int) *ListNode {

	pseudo := &ListNode{Val: 0}
	pseudo.Next = head
	cur := pseudo
	for cur != nil && cur.Next != nil {
		if cur.Next.Val == val {
			cur.Next = cur.Next.Next
		} else {
			cur = cur.Next
		}

	}
	return pseudo.Next
}
