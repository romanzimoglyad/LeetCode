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
		Val: 4,
	}
	st5 := &ListNode{
		Val: 5,
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
	test1 := &ListNode{
		Val: 2,
	}
	test2 := &ListNode{
		Val: 7,
	}

	test3 := &ListNode{
		Val: 26,
	}
	test1.Next = test2
	test2.Next = test3
	res := mergeTwoLists(st1, test1)
	fmt.Println(res)
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	resp := &ListNode{}

	/*pseudo1 := &ListNode{
		Val:  0,
		Next: list1,
	}
	pseudo2 := &ListNode{
		Val:  0,
		Next: list2,
	}*/
	cur1 := list1
	cur2 := list2
	curResp := resp
	for cur1 != nil || cur2 != nil {
		if cur1 == nil {
			curResp.Next = cur2
			return resp.Next
		}
		if cur2 == nil {
			curResp.Next = cur1
			return resp.Next
		}

		if cur1.Val <= cur2.Val {
			curResp.Next = cur1
			curResp = curResp.Next
			cur1 = cur1.Next
		} else {
			curResp.Next = cur2
			curResp = curResp.Next
			cur2 = cur2.Next
		}
	}
	return resp.Next
}
