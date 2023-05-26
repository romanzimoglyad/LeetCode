package main

import "fmt"

func main() {
	obj := Constructor()
	param_1 := obj.Get(1)
	fmt.Println(param_1)
	obj.AddAtHead(1)

	obj.DeleteAtIndex(0)

}

type Node struct {
	val  int
	next *Node
}

type MyLinkedList struct {
	head   *Node
	length int
}

func Constructor() MyLinkedList {
	return MyLinkedList{
		head:   nil,
		length: 0,
	}
}

func (m *MyLinkedList) Get(index int) int {
	if index >= m.length {
		return -1
	}
	cur := m.head
	for i := 0; i < index; i++ {
		cur = cur.next
	}
	return cur.val

}

func (m *MyLinkedList) AddAtHead(val int) {
	newNode := &Node{
		val:  val,
		next: nil,
	}
	if m.length == 0 {
		m.head = newNode
		m.length++
	} else {
		newNode.next = m.head
		m.head = newNode
		m.length++
	}
}

func (m *MyLinkedList) AddAtTail(val int) {
	newNode := &Node{
		val: val,
	}
	if m.length == 0 {
		m.head = newNode
		m.length++
		return
	}
	cur := m.head
	for cur.next != nil {
		cur = cur.next
	}
	cur.next = newNode
	m.length++

}

func (m *MyLinkedList) AddAtIndex(index int, val int) {
	if index > m.length {
		return
	}
	if index == m.length {
		m.AddAtTail(val)
		return
	}
	if index == 0 {
		m.AddAtHead(val)
		return
	}
	cur := m.head
	for i := 0; i < index-1; i++ {
		cur = cur.next
	}
	newNode := &Node{
		val:  val,
		next: cur.next,
	}
	cur.next = newNode
	m.length++
}

func (m *MyLinkedList) DeleteAtIndex(index int) {
	if index >= m.length {
		return
	}
	if index == 0 {
		m.head = m.head.next
		m.length--
		return
	}

	cur := m.head
	for i := 0; i < index-1; i++ {
		cur = cur.next
	}
	cur.next = cur.next.next
	m.length--
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
