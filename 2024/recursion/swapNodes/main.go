package main

import "fmt"

func main() {
	l1 := &ListNode{
		Val:  1,
		Next: nil,
	}
	l2 := &ListNode{
		Val:  2,
		Next: nil,
	}
	l3 := &ListNode{
		Val:  3,
		Next: nil,
	}
	l4 := &ListNode{
		Val:  4,
		Next: nil,
	}
	l1.Next = l2
	l2.Next = l3

	l3.Next = l4
	resp := swapPairs(l1)
	fmt.Println(resp)

}

type ListNode struct {
	Val  int
	Next *ListNode
}

func swapPairs(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	tail := head.Next.Next

	head.Next.Next = head
	resp := head.Next
	head.Next = swapPairs(tail)
	return resp
}
