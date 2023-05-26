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

	/*	st5 := &ListNode{
			Val: 1,
		}
		st4 := &ListNode{
			Val:  4,
			Next: st5,
		}

		st5.Next = st2*/

	res := reverseList(st1)
	fmt.Println(res)
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	prev := reverseList(head.Next)
	head.Next.Next = head
	head.Next = nil
	return prev
}
