package main

import "fmt"

func main() {
	test1 := &ListNode{
		Val: 1,
	}
	test2 := &ListNode{
		Val: 2,
	}

	test3 := &ListNode{
		Val: 3,
	}
	test1.Next = test2
	test2.Next = test3
	test3.Next = test2

	_ = &ListNode{
		Val: 1,
	}
	t1 := hasCycle(test1)
	fmt.Println(t1)
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func hasCycle(head *ListNode) bool {
	if head == nil {
		return false
	}
	fast := head
	slow := head

	for fast.Next != nil && fast.Next.Next != nil {

		fast = fast.Next.Next
		slow = slow.Next
		if fast == slow {
			return true
		}
	}
	return false
}
