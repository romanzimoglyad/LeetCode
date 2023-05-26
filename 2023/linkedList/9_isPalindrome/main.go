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
	st4 := &ListNode{
		Val: 2,
	}
	st5 := &ListNode{
		Val: 1,
	}
	st2.Next = st3
	st3.Next = st4
	st4.Next = st5
	/*	st5 := &ListNode{
			Val: 1,
		}
		st4 := &ListNode{
			Val:  4,
			Next: st5,
		}

		st5.Next = st2*/

	res := isPalindrome(st1)
	fmt.Println(res)
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func isPalindrome(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return true
	}
	i := 0
	cur := head
	for cur != nil {
		cur = cur.Next
		i++
	}
	newLen := i / 2
	newCur := head
	for k := 0; k < newLen; k++ {
		newCur = newCur.Next
	}
	reversedEnd := reverseList(newCur)

	againCur := head
	for k := 0; k < newLen; k++ {
		if againCur.Val != reversedEnd.Val {
			return false
		}
		againCur = againCur.Next
		reversedEnd = reversedEnd.Next
	}
	return true
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
