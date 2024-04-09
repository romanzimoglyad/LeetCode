package main

import "fmt"

func main() {
	n4 := &ListNode{
		Val: 4,
	}
	n3 := &ListNode{
		Val: 3,
	}
	n2 := &ListNode{
		Val: 2,
	}
	n1 := &ListNode{
		Val: 1,
	}
	n4.Next = n2
	n2.Next = n1
	n1.Next = n3
	resp := insertionSortList(n4)
	fmt.Println(resp)
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func insertionSortList(head *ListNode) *ListNode {
	cur := head.Next
	dummyResp := &ListNode{}
	dummyResp.Next = &ListNode{Val: head.Val}

	for cur != nil {
		tmpHead := dummyResp
		for tmpHead.Next != nil && tmpHead.Next.Val < cur.Val {
			tmpHead = tmpHead.Next
		}
		next := tmpHead.Next
		tmpHead.Next = &ListNode{Val: cur.Val}
		tmpHead.Next.Next = next
		cur = cur.Next
	}

	return dummyResp.Next
}

//for prev != nil {
//if cur.Val < prev.Val {
//tmpNext := cur.Next
//cur.Next = prev
//prev.Next = tmpNext
//}
//}
//cur = cur.Next
