package main

import (
	"fmt"
	"sync"
)

func main() {
	c := Constructor(3)
	//c.Put(1, 1)
	//c.Put(2, 2)
	//fmt.Println(c.Get(1))
	//c.Put(3, 3)
	//fmt.Println(c.Get(2))
	//c.Put(4, 4)
	//fmt.Println(c.Get(1))
	//fmt.Println(c.Get(3))
	//fmt.Println(c.Get(4))

	//c.Put(2, 1)
	//fmt.Println(c.Get(2))
	//c.Put(3, 2)
	//
	//fmt.Println(c.Get(2))
	//fmt.Println(c.Get(3))
	//c.Put(2, 1)
	//c.Put(2, 2)
	//fmt.Println(c.Get(2))
	//c.Put(1, 1)
	//c.Put(4, 1)
	//fmt.Println(c.Get(2))
	c.Put(1, 1)
	c.Put(2, 2)
	c.Put(3, 3)
	c.Put(4, 4)
	c.Put(4, 44)
	fmt.Println(c.Get(4))
	fmt.Println(c.Get(3))
	fmt.Println(c.Get(2))
	fmt.Println(c.Get(1))
	c.Put(5, 5)
	fmt.Println(c.Get(1))
	fmt.Println(c.Get(2))
	fmt.Println(c.Get(3))
	fmt.Println(c.Get(4))
	fmt.Println(c.Get(5))

}

type LRUCache struct {
	d     DoubleLinkedList
	m     map[int]*Node
	cap   int
	size  int
	mutex sync.RWMutex
}

func Constructor(capacity int) LRUCache {

	return LRUCache{m: make(map[int]*Node), cap: capacity}
}

func (this *LRUCache) Get(key int) int {

	this.mutex.RLock()
	defer this.mutex.RUnlock()
	if this.size == 0 {
		return -1
	}
	var val *Node
	ok := false
	if val, ok = this.m[key]; ok {
		if this.size > 1 {
			this.d.delete(val)
			this.d.addFront(val)
		}
	} else {
		return -1
	}

	return val.value
}

func (this *LRUCache) Put(key int, value int) {
	this.mutex.Lock()
	defer this.mutex.Unlock()

	if node, ok := this.m[key]; ok {
		node.value = value
		this.m[key] = node
		this.d.delete(node)
		this.d.addFront(node)
	} else {
		newNode := &Node{
			key:   key,
			value: value,
		}
		this.m[key] = newNode
		if this.size == this.cap {
			key := this.d.deleteLast()
			delete(this.m, key)
			this.size--
		}
		this.size++
		this.d.addFront(newNode)
	}

}

type Node struct {
	key   int
	value int
	next  *Node
	prev  *Node
}

type DoubleLinkedList struct {
	head *Node
	tail *Node
}

func (d *DoubleLinkedList) delete(val *Node) {

	if d.tail == d.head && d.tail == val {
		d.tail = nil
		d.head = nil
		return
	}
	if val == d.head {
		d.head = d.head.next
		d.head.prev = nil
	} else if val == d.tail {
		d.tail = d.tail.prev
		val.prev.next = nil
	} else {
		val.prev.next = val.next
		val.next.prev = val.prev
	}
}

func (d *DoubleLinkedList) deleteLast() int {
	if d.tail == nil {
		return -1
	}
	ind := d.tail.key
	if d.tail == d.head {
		d.tail = nil
		d.head = nil
		return ind
	}
	d.tail.prev.next = nil
	d.tail = d.tail.prev
	return ind

}

func (d *DoubleLinkedList) append(newNode *Node) {

	if d.head == nil {
		d.head = newNode
		d.tail = newNode
		return
	}
	newNode.prev = d.tail
	d.tail.next = newNode
	d.tail = newNode
}

func (d *DoubleLinkedList) getTail() *Node {
	return d.tail
}

func (d *DoubleLinkedList) addFront(newNode *Node) {
	if d.head == nil {
		d.append(newNode)
	} else {
		newNode.next = d.head

		d.head.prev = newNode
		d.head = newNode
	}

}

func (d *DoubleLinkedList) getFront() *Node {
	return d.head
}

/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */
