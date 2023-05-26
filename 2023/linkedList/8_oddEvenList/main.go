package main

import "fmt"

func main() {

	st2 := &ListNode{
		Val: 1,
	}
	st1 := &ListNode{
		Val:  2,
		Next: st2,
	}
	st3 := &ListNode{
		Val: 3,
	}
	st4 := &ListNode{
		Val: 5,
	}
	st2.Next = st3
	st3.Next = st4
	/*	st5 := &ListNode{
			Val: 1,
		}
		st4 := &ListNode{
			Val:  4,
			Next: st5,
		}

		st5.Next = st2*/

	res := oddEvenList(st1)
	fmt.Println(res)
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func oddEvenList(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	pseudoFirstOdd := &ListNode{
		Val: 0,
	}
	var LastOdd *ListNode = pseudoFirstOdd
	pseudoFirstEven := &ListNode{
		Val: 0,
	}
	var lastEven *ListNode = pseudoFirstEven
	cur := head
	i := 1
	for cur != nil {

		if i%2 == 0 {
			lastEven.Next = cur
			lastEven = lastEven.Next
		} else if i%2 == 1 {
			LastOdd.Next = cur
			LastOdd = LastOdd.Next
		}
		if cur.Next == nil {
			lastEven.Next = nil
		}
		cur = cur.Next
		i++
	}

	LastOdd.Next = pseudoFirstEven.Next

	return pseudoFirstOdd.Next
}
