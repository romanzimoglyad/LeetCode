package main

import "fmt"

func main() {

	st2 := &ListNode{
		Val: 4,
	}
	st1 := &ListNode{
		Val:  2,
		Next: st2,
	}
	st3 := &ListNode{
		Val: 3,
	}
	st2.Next = st3
	/*st4 := &ListNode{
		Val: 4,
	}
	st5 := &ListNode{
		Val: 5,
	}

	st3.Next = st4
	st4.Next = st5*/
	/*	st5 := &ListNode{
			Val: 1,
		}
		st4 := &ListNode{
			Val:  4,
			Next: st5,
		}

		st5.Next = st2*/
	test1 := &ListNode{
		Val: 5,
	}
	test2 := &ListNode{
		Val: 6,
	}

	test3 := &ListNode{
		Val: 9,
	}
	test1.Next = test2
	test2.Next = test3
	res := addTwoNumbers(st1, test1)
	fmt.Println(res)
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	cur1 := l1
	cur2 := l2
	pseudo := &ListNode{}
	cur := pseudo
	remain := false
	for cur1 != nil || cur2 != nil {
		cur2Val := 0
		cur1Val := 0
		if cur2 == nil {
			cur2Val = 0
		} else {
			cur2Val = cur2.Val
			cur2 = cur2.Next
		}
		if cur1 == nil {
			cur1Val = 0
		} else {
			cur1Val = cur1.Val
			cur1 = cur1.Next
		}
		val := 0
		if remain {
			val = cur1Val + cur2Val + 1
			remain = false
		} else {
			val = cur1Val + cur2Val
		}
		value := val % 10
		remain = val >= 10
		cur.Next = &ListNode{
			Val: value,
		}
		cur = cur.Next
	}
	if remain {
		cur.Next = &ListNode{
			Val: 1,
		}
	}
	return pseudo.Next
}
