package main

import "fmt"

func main() {

	test1 := &ListNode{
		Val: 1,
	}
	test2 := &ListNode{
		Val: 2,
	}

	test3 := &ListNode{
		Val: 3,
	}
	test4 := &ListNode{
		Val: 4,
	}
	test1.Next = test2
	test2.Next = test3
	test3.Next = test4
	test4.Next = test2
	_ = &ListNode{
		Val: 1,
	}
	t1 := detectCycle(test1)
	fmt.Println(t1)
}

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 *



}
*/
func detectCycle(head *ListNode) *ListNode {
	inter := detectCycle1(head)
	if inter == nil {
		return nil
	}
	ptr1 := inter
	ptr2 := head
	for ptr1 != ptr2 {
		ptr1 = ptr1.Next
		ptr2 = ptr2.Next
	}

	return ptr1
}
func detectCycle1(head *ListNode) *ListNode {

	if head == nil {
		return nil
	}
	curNodefast := head
	curNodeSlow := head
	for curNodefast.Next != nil && curNodefast.Next.Next != nil {
		curNodeSlow = curNodeSlow.Next
		curNodefast = curNodefast.Next.Next

		if curNodeSlow == curNodefast {
			return curNodefast
		}
	}
	return nil
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
