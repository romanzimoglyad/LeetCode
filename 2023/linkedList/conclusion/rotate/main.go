package main

import "fmt"

func main() {
	/*st4 := &Node{
		Val: 4,
	}
	*/
	/*
		st3 := &Node{
			Val: 3,
		}*/

	/*	st2 := &Node{
		Val: 2,
	}*/
	st1 := &ListNode{

		Val: 1,
	}
	st2 := &ListNode{

		Val: 2,
	}
	st3 := &ListNode{

		Val: 3,
	}
	st4 := &ListNode{

		Val: 4,
	}
	st5 := &ListNode{

		Val: 5,
	}
	st1.Next = st2
	st2.Next = st3
	st3.Next = st4
	st4.Next = st5
	//	st1.Next = st2
	//st2.Next = st3
	//	st2.Random = st1
	//st3.Random = st1
	//st4.Next = st1

	res := rotateRight(st1, 10)
	fmt.Println(res)
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func rotateRight(head *ListNode, k int) *ListNode {
	if k == 0 {
		return head
	}
	if head == nil {
		return nil
	}
	length := 0
	cur := head
	for cur != nil {
		length++
		cur = cur.Next
	}
	pp := k % length
	if pp == 0 || length == 1 {
		return head
	}
	cur = head
	t := length - k%length - 1
	for i := 0; i < t; i++ {
		cur = cur.Next
	}
	resp := cur.Next
	oldCur := cur
	for cur.Next != nil {
		cur = cur.Next
	}
	oldCur.Next = nil
	cur.Next = head

	return resp

}
func goDeap(head *ListNode) {

}
