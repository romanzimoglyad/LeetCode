package main

import "fmt"

func main() {
	list := Constructor()
	list.AddAtTail(3)
	list.AddAtHead(1)
	list.AddAtIndex(1, 2)
	v1 := list.Get(0)

	v2 := list.Get(1)
	v3 := list.Get(2)
	fmt.Println(v1, v2, v3)
	list.DeleteAtIndex(1)
	v11 := list.Get(0)

	v22 := list.Get(1)
	v33 := list.Get(2)
	fmt.Println(v11, v22, v33)
	list.AddAtIndex(45, 2)
}

type MyLinkedList struct {
	length int
	head   *Node
}

type Node struct {
	val  int
	next *Node
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
		if curNode.next == nil {
			return -1
		}
		curNode = curNode.next
		i++
	}
	return curNode.val
}

/** Add a node of value val before the first element of the linked list. After the insertion, the new node will be the first node of the linked list. */
func (this *MyLinkedList) AddAtHead(val int) {
	cur := this.head

	newNode := &Node{
		val:  val,
		next: cur,
	}

	this.head = newNode
	this.length++
}

/** Append a node of value val to the last element of the linked list. */
func (this *MyLinkedList) AddAtTail(val int) {
	var curNode *Node
	curNode = this.head
	newNode := &Node{
		val: val,
	}
	if curNode == nil {
		this.head = newNode
		this.length++
		return
	}
	for curNode.next != nil {
		curNode = curNode.next
	}
	curNode.next = newNode
	this.length++
}

/** Add a node of value val before the index-th node in the linked list. If index equals to the length of linked list, the node will be appended to the end of linked list. If index is greater than the length, the node will not be inserted. */
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
	var curNode *Node
	curNode = this.head
	newNode := &Node{
		val: val,
	}
	if curNode == nil {
		this.head = newNode
		return
	}
	i := 0
	for i < index-1 {
		curNode = curNode.next
		i++
	}
	tmp := curNode.next
	newNode.next = tmp
	curNode.next = newNode
	this.length++
}

/** Delete the index-th node in the linked list, if the index is valid. */
func (this *MyLinkedList) DeleteAtIndex(index int) {
	if index >= this.length {
		return
	}
	if index == 0 {
		this.head = this.head.next
		this.length--
		return
	}
	i := 0
	curNode := this.head
	for i < index-1 {
		curNode = curNode.next
		i++
	}
	curNode.next = curNode.next.next
	this.length--
}

/**
 * Your MyLinkedList object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Get(index);
 * obj.AddAtHead(val);
 * obj.AddAtTail(val);
 * obj.AddAtIndex(index,val);
 * obj.DeleteAtIndex(index);
 */
