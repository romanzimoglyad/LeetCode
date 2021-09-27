package main

import (
	"fmt"
)

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

	test1.Next = test2

	test2.Next = test3
	r := rotateRight(test1, 1)
	fmt.Println(r)
}

func rotateRight(head *ListNode, k int) *ListNode {

	if head == nil {
		return nil
	}
	if head.Next == nil {
		return head
	}
	length := 1
	tmp := head
	for tmp.Next != nil {
		tmp = tmp.Next
		length++
	}
	p := k % length

	cur := head
	for i := 0; i < p; i++ {
		cur = rotateOnes(cur)
	}
	return cur
}

func rotateOnes(head *ListNode) *ListNode {
	cur := head
	prev := head
	for cur.Next != nil {
		prev = cur
		cur = cur.Next
	}
	prev.Next = nil
	cur.Next = head
	return cur
}

type MyLinkedList struct {
	length int
	head   *ListNode
}

type ListNode struct {
	Val  int
	Next *ListNode

	Prev *ListNode
}

/** Initialize your data structure here. */
func Constructor() MyLinkedList {
	return MyLinkedList{head: nil}
}

/** Get the value of the index-th node in the linked list. If the index is invalid, return -1. */
func (this *MyLinkedList) Get(index int) int {
	if index >= this.length {
		return -1
	}
	i := 0
	cur := this.head
	for i < index {
		cur = cur.Next
		i++
	}
	return cur.Val

}

/** Add a node of value Val before the first element of the linked list. After the insertion, the new node will be the first node of the linked list. */
func (this *MyLinkedList) AddAtHead(val int) {
	if this.head == nil {
		this.head = &ListNode{
			Val: val,
		}
		this.length++
		return
	}
	newHead := &ListNode{
		Val:  val,
		Next: this.head,
	}
	this.head.Prev = newHead
	this.head = newHead
	this.length++
}

/** Append a node of value Val to the last element of the linked list. */
func (this *MyLinkedList) AddAtTail(val int) {
	cur := this.head
	if cur == nil {
		this.head = &ListNode{
			Val: val,
		}
		this.length++
		return
	}
	for cur.Next != nil {
		cur = cur.Next
	}
	newHead := &ListNode{
		Val:  val,
		Prev: cur,
	}
	cur.Next = newHead
	this.length++
}

/** Add a node of value Val before the index-th node in the linked list. If index equals to the length of linked list, the node will be appended to the end of linked list. If index is greater than the length, the node will not be inserted. */
func (this *MyLinkedList) AddAtIndex(index int, val int) {
	if index >= this.length {
		return
	}
	if index == 0 {
		this.AddAtHead(val)
		return
	}
	if index == this.length {
		this.AddAtTail(val)
		return
	}
	i := 0
	cur := this.head
	for i < index-1 {
		cur = cur.Next
	}

	newNode := &ListNode{
		Val:  val,
		Prev: cur,
		Next: cur.Next,
	}
	cur.Next = newNode
	newNode.Next.Prev = newNode
	this.length++
}

/** Delete the index-th node in the linked list, if the index is valid. */
func (this *MyLinkedList) DeleteAtIndex(index int) {
	if index >= this.length {
		return
	}
	if this.length == 1 {
		this.head = nil
		this.length--
		return
	}
	if index == 0 {
		this.head = this.head.Next
		this.head.Prev = nil
		this.length--
		return
	}
	if index == this.length-1 {
		cur := this.head
		for cur.Next != nil {
			cur = cur.Next
		}
		cur.Prev.Next = nil
		this.length--
		return
	}
	i := 0
	cur := this.head

	for i < index-1 {
		cur = cur.Next
	}
	cur.Next = cur.Next.Next
	cur.Next.Prev = cur
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
