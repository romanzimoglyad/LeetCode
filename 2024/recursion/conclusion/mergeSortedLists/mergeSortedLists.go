package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	if list1 == nil {
		return list2
	}
	if list2 == nil {
		return list1
	}
	next1 := list1.Next
	next2 := list2.Next

	if list2.Val >= list1.Val {
		list1.Next = mergeTwoLists(next1, list2)
		return list1
	}
	list2.Next = mergeTwoLists(list1, next2)
	return list2

}
