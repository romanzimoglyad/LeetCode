package main

func main() {
}

type ListNode struct {
	Val  int
	Next *ListNode
}

// 1 2 3
func reverseList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	resp := reverseList(head.Next)
	head.Next.Next = head
	head.Next = nil
	return resp
}
