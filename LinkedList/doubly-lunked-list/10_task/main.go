package main

import "fmt"

func main() {
	list := Constructor()

	list.AddAtTail(4)
	list.AddAtHead(1)
	list.AddAtIndex(1, 3)
	list.AddAtIndex(1, 2)
	v1 := list.Get(0)

	v2 := list.Get(1)
	v3 := list.Get(2)
	v4 := list.Get(3)
	fmt.Println(v1, v2, v3, v4)
	list.DeleteAtIndex(1)
	list.DeleteAtIndex(55)
	list.DeleteAtIndex(0)
	v11 := list.Get(0)
	v22 := list.Get(1)
	v33 := list.Get(2)
	v44 := list.Get(3)
	fmt.Println(v11, v22, v33, v44)
	list.AddAtIndex(45, 2)

	newList := Constructor()
	newList.AddAtHead(1)
	newList.DeleteAtIndex(0)

}

type MyLinkedList struct {
	length int
	head   *Node
}

type Node struct {
	val  int
	next *Node
	prev *Node
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
		cur = cur.next
		i++
	}
	return cur.val

}

/** Add a node of value val before the first element of the linked list. After the insertion, the new node will be the first node of the linked list. */
func (this *MyLinkedList) AddAtHead(val int) {
	if this.head == nil {
		this.head = &Node{
			val: val,
		}
		this.length++
		return
	}
	newHead := &Node{
		val:  val,
		next: this.head,
	}
	this.head.prev = newHead
	this.head = newHead
	this.length++
}

/** Append a node of value val to the last element of the linked list. */
func (this *MyLinkedList) AddAtTail(val int) {
	cur := this.head
	if cur == nil {
		this.head = &Node{
			val: val,
		}
		this.length++
		return
	}
	for cur.next != nil {
		cur = cur.next
	}
	newHead := &Node{
		val:  val,
		prev: cur,
	}
	cur.next = newHead
	this.length++
}

/** Add a node of value val before the index-th node in the linked list. If index equals to the length of linked list, the node will be appended to the end of linked list. If index is greater than the length, the node will not be inserted. */
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
		cur = cur.next
	}

	newNode := &Node{
		val:  val,
		prev: cur,
		next: cur.next,
	}
	cur.next = newNode
	newNode.next.prev = newNode
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
		this.head = this.head.next
		this.head.prev = nil
		this.length--
		return
	}
	if index == this.length-1 {
		cur := this.head
		for cur.next != nil {
			cur = cur.next
		}
		cur.prev.next = nil
		this.length--
		return
	}
	i := 0
	cur := this.head

	for i < index-1 {
		cur = cur.next
	}
	cur.next = cur.next.next
	cur.next.prev = cur
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
