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
	test11 := &ListNode{
		Val: 1,
	}
	test22 := &ListNode{
		Val: 2,
	}
	/*test33 := &ListNode{
		Val: 2,
	}*/
	test1.Next = test2
	test2.Next = test3

	test11.Next = test22
	test22.Next = test3

	_ = &ListNode{
		Val: 1,
	}

	t1 := reverseList(test1)
	fmt.Println(t1)
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseList(head *ListNode) *ListNode {

	if head.Next == nil || head == nil {
		return head
	}

	res := reverseList(head.Next)
	head.Next.Next = head
	head.Next = nil
	return res
}
