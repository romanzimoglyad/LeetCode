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
	st2.Next = st3

	st5 := &ListNode{
		Val: 1,
	}
	st4 := &ListNode{
		Val:  4,
		Next: st5,
	}

	st5.Next = st2

	res := getIntersectionNode(st1, st4)
	fmt.Println(res)
}

type ListNode struct {
	Val  int
	Next *ListNode
}

/**
 * Definition for singly-linked list.
 *
 */
func getIntersectionNode(headA, headB *ListNode) *ListNode {
	m := make(map[*ListNode]struct{})
	cur1 := headA
	for cur1 != nil {
		if _, ok := m[cur1]; !ok {
			m[cur1] = struct{}{}
		}
		cur1 = cur1.Next
	}
	cur2 := headB
	for cur2 != nil {
		if _, ok := m[cur2]; ok {
			return cur2
		}
		cur2 = cur2.Next
	}
	return nil
}
