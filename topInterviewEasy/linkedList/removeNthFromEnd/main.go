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

	t1 := removeNthFromEnd(test1, 2)
	fmt.Println(t1)
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	i := 0
	tmpHead := &ListNode{
		Val:  0,
		Next: head,
	}
	cur := tmpHead
	cur2 := tmpHead
	for cur.Next != nil {
		cur = cur.Next
		i++
		if i >= n+1 {
			cur2 = cur2.Next
		}
	}

	cur2.Next = cur2.Next.Next
	return tmpHead.Next
}
