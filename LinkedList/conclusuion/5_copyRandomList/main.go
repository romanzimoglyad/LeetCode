package main

import (
	"fmt"
)

func main() {
	test1 := &Node{
		Val: 1,
	}
	test2 := &Node{
		Val: 2,
	}

	test3 := &Node{
		Val: 3,
	}

	test1.Next = test2
	test1.Random = test2
	test2.Next = test3

	r := copyRandomList(test1)
	fmt.Println(r)
}

func copyRandomList(head *Node) *Node {
	if head == nil {
		return nil
	}
	cur := head
	newHead := &Node{}
	var newCur = newHead
	mapping := map[*Node]*Node{}
	for cur != nil && newCur != nil {
		if existNode, ok := mapping[cur]; !ok {
			newCur.Val = cur.Val
			mapping[cur] = newCur
		} else {
			newCur = existNode
		}
		if existNode, ok := mapping[cur.Next]; !ok && cur.Next != nil {
			newCurNext := &Node{
				Val: cur.Next.Val,
			}
			newCur.Next = newCurNext
			mapping[cur.Next] = newCurNext
		} else {
			newCur.Next = existNode
		}
		if existRandomNode, ok := mapping[cur.Random]; !ok && cur.Random != nil {
			newCurRandomNext := &Node{
				Val: cur.Random.Val,
			}
			newCur.Random = newCurRandomNext
			mapping[cur.Random] = newCurRandomNext
		} else {
			newCur.Random = existRandomNode
		}
		cur = cur.Next
		newCur = newCur.Next
	}
	return newHead
}

type MyLinkedList struct {
	length int
	head   *Node
}

type Node struct {
	Val    int
	Next   *Node
	Random *Node
	Prev   *Node
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
		this.head = &Node{
			Val: val,
		}
		this.length++
		return
	}
	newHead := &Node{
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
		this.head = &Node{
			Val: val,
		}
		this.length++
		return
	}
	for cur.Next != nil {
		cur = cur.Next
	}
	newHead := &Node{
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

	newNode := &Node{
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
