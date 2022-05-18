package main

import "fmt"

func main() {
	// t := hasCycle(list.head)
	test1 := &ListNode{
		Val: 1,
	}
	test2 := &ListNode{
		Val: 2,
	}

	// test5 := &ListNode{
	//	Val: 5,
	// }
	/*test33 := &ListNode{
		Val: 2,
	}*/
	test1.Next = test2

	// test4.Next = test5
	_ = &ListNode{
		Val: 1,
	}

	t1 := isPalindrome(test1)
	fmt.Println(t1)
}

type ListNode struct {
	Val  int
	Next *ListNode
}

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func isPalindrome(head *ListNode) bool {
	reverse := reverse(head)

	cur1 := head
	cur2 := reverse

	for cur1 != nil && cur2 != nil {
		if cur1.Val != cur2.Val {
			return false
		}
		cur1 = cur1.Next
		cur2 = cur2.Next
	}
	if cur1 != nil || cur2 != nil {
		return false
	}
	return true
}

func reverse(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	var prev *ListNode
	current := head

	for current != nil {
		nxt := current.Next
		current.Next = prev
		prev = current
		current = nxt
	}
	return prev
}
