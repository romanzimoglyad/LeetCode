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

	test3 := &ListNode{
		Val: 3,
	}
	test11 := &ListNode{
		Val: 4,
	}
	test22 := &ListNode{
		Val: 5,
	}

	test1.Next = test2
	test2.Next = test3
	test3.Next = test11
	test11.Next = test22

	_ = &ListNode{
		Val: 1,
	}

	t1 := reverseListt(test1)
	fmt.Println(t1)
}

func reverseList3(head *ListNode) (*ListNode, *ListNode) {
	if head.Next == nil {
		return head,head
	}
	r1, t1 := reverseList3(head.Next)
	t1.Next = head
	head.Next = nil
	return r1, head
}

func reverseListt(head *ListNode) *ListNode {
	newHead, _ := reverseList3(head)
	return newHead
}

func reverseList2(head *ListNode) *ListNode {
	if head.Next == nil {
		return head
	}
	r1 := reverseList2(head.Next)

	for tail := r1;tail.Next!=nil; tail=tail.Next {
	}
	r1.Next = head
	head.Next = nil
	return r1
}


/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func reverseList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	p := reverseList(head.Next)
	head.Next.Next = head
	head.Next = nil
	return p
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	dum := &ListNode{}
	dum.Next = head
	cur1 := dum
	cur2 := dum
	for i := 0; i <= n; i++ {
		cur1 = cur1.Next
	}
	for cur1 != nil {
		cur1 = cur1.Next
		cur2 = cur2.Next
	}
	cur2.Next = cur2.Next.Next
	return dum.Next
}

type MyLinkedList struct {
	length int
	head   *ListNode
}

type ListNode struct {
	Val  int
	Next *ListNode
}

/** Initialize your data structure here. */
func Constructor() MyLinkedList {
	return MyLinkedList{head: nil}
}

/** Get the value of the index-th node in the linked list. If the index is invalid, return -1. */
func (this *MyLinkedList) Get(index int) int {
	i := 0
	if this.head == nil {
		return -1
	}
	if index >= this.length {
		return -1
	}

	curNode := this.head
	for i < index {
		if curNode.Next == nil {
			return -1
		}
		curNode = curNode.Next
		i++
	}
	return curNode.Val
}

/** Add a node of value Val before the first element of the linked list. After the insertion, the new node will be the first node of the linked list. */
func (this *MyLinkedList) AddAtHead(val int) {
	cur := this.head

	newNode := &ListNode{
		Val:  val,
		Next: cur,
	}

	this.head = newNode
	this.length++
}

/** Append a node of value Val to the last element of the linked list. */
func (this *MyLinkedList) AddAtTail(val int) {
	var curNode *ListNode
	curNode = this.head
	newNode := &ListNode{
		Val: val,
	}
	if curNode == nil {
		this.head = newNode
		this.length++
		return
	}
	for curNode.Next != nil {
		curNode = curNode.Next
	}
	curNode.Next = newNode
	this.length++
}

/** Add a node of value Val before the index-th node in the linked list. If index equals to the length of linked list, the node will be appended to the end of linked list. If index is greater than the length, the node will not be inserted. */
func (this *MyLinkedList) AddAtIndex(index int, val int) {
	if index == 0 {
		this.AddAtHead(val)
		return
	}
	if index == this.length {
		this.AddAtTail(val)
		return
	}
	if index > this.length {
		return
	}
	var curNode *ListNode
	curNode = this.head
	newNode := &ListNode{
		Val: val,
	}
	if curNode == nil {
		this.head = newNode
		return
	}
	i := 0
	for i < index-1 {
		curNode = curNode.Next
		i++
	}
	tmp := curNode.Next
	newNode.Next = tmp
	curNode.Next = newNode
	this.length++
}

/** Delete the index-th node in the linked list, if the index is valid. */
func (this *MyLinkedList) DeleteAtIndex(index int) {
	if index >= this.length {
		return
	}
	if index == 0 {
		this.head = this.head.Next
		this.length--
		return
	}
	i := 0
	curNode := this.head
	for i < index-1 {
		curNode = curNode.Next
		i++
	}
	curNode.Next = curNode.Next.Next
	this.length--
}

/**
 * Your MyLinkedList object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Get(index);
 * obj.AddAtHead(Val);
 * obj.AddAtTail(Val);
 * obj.AddAtIndex(index,Val);
 * obj.DeleteAtIndex(index);
 */
