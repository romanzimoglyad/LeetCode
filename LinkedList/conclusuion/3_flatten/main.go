package main

import "fmt"

func main() {
	/*	list := Constructor()

		list.AddAtTail(4)
		list.AddAtHead(1)
		list.AddAtIndex(1, 3)
		list.AddAtIndex(1, 2)
		v1 := list.Get(0)*/
	/*
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
		newList.DeleteAtIndex(0)*/

	v1 := &Node{
		Val:   1,
		Next:  nil,
		Prev:  nil,
		Child: nil,
	}
	child1 := &Node{
		Val:   2,
		Next:  nil,
		Prev:  nil,
		Child: nil,
	}
	child2 := &Node{
		Val:   33,
		Next:  nil,
		Prev:  nil,
		Child: nil,
	}
	child2Next := &Node{
		Val:   333,
		Next:  nil,
		Prev:  nil,
		Child: nil,
	}
	childNext := &Node{
		Val:   222,
		Next:  nil,
		Prev:  nil,
		Child: nil,
	}

	v1Next := &Node{
		Val:   11,
		Next:  nil,
		Prev:  nil,
		Child: nil,
	}

	v1.Child = child1
	v1.Next = v1Next
	child1.Child = child2
	child1.Next = childNext
	child2.Next = child2Next
	r := flatten(nil)
	fmt.Println(r)
}

func flatten(root *Node) *Node {
	if root == nil {
		return nil
	}
	dum := &Node{Next: &Node{}}
	handleDeep(nil, dum, root)

	return dum
}

// 1  2 33 333 222 11
func handleDeep(prev *Node, resp *Node, node *Node) *Node {
	resp.Val = node.Val
	resp.Prev = prev
	if node == nil {
		return nil
	}
	if node.Child != nil {
		resp.Next = &Node{}
		resp = handleDeep(resp, resp.Next, node.Child)
	}

	if node.Next != nil {
		resp.Next = &Node{}
		resp = handleDeep(resp, resp.Next, node.Next)
	}
	return resp
}

type MyLinkedList struct {
	length int
	head   *Node
}

type Node struct {
	Val   int
	Next  *Node
	Prev  *Node
	Child *Node
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
