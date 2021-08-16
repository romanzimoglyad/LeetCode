package main

import "fmt"

func main() {
	list := Constructor()
	list.AddAtTail(3)
	list.AddAtHead(1)

	//t := hasCycle(list.head)
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
		Val: 1,
	}
	test22 := &ListNode{
		Val: 2,
	}

	test1.Next = test2
	test2.Next = test3

	test11.Next = test22
	test22.Next = test3

	_ = &ListNode{
		Val: 1,
	}

	t1 := getIntersectionNode(test1, test11)
	fmt.Println(t1)
}

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func getIntersectionNode(headA, headB *ListNode) *ListNode {
	mapA := []*ListNode{}
	mapB := make(map[*ListNode]struct{})
	curA := headA
	for curA.Next != nil {
		mapA = append(mapA, curA)
		curA = curA.Next
	}
	mapA = append(mapA, curA)
	curB := headB
	for curB.Next != nil {

		mapB[curB] = struct{}{}

		curB = curB.Next
	}
	mapB[curB] = struct{}{}

	for _, elem := range mapA {
		if _, ok := mapB[elem]; ok {
			return elem
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
