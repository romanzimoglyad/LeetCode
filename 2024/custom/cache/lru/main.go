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
	m     map[int]int
	cap   int
	size  int
	mutex sync.RWMutex
}

func Constructor(capacity int) LRUCache {

	return LRUCache{m: make(map[int]int), cap: capacity}
}

func (this *LRUCache) Get(key int) int {

	this.mutex.RLock()
	defer this.mutex.RUnlock()
	if this.size == 0 {
		return -1
	}
	val := 0
	ok := false
	if val, ok = this.m[key]; ok {
		if this.size > 1 {
			this.d.delete(key)
			this.d.addFront(key)
		}
	} else {
		return -1
	}

	return val
}

func (this *LRUCache) Put(key int, value int) {
	this.mutex.Lock()
	defer this.mutex.Unlock()

	if _, ok := this.m[key]; ok {
		this.m[key] = value
		this.d.delete(key)
		this.d.addFront(key)
	} else {
		this.m[key] = value
		if this.size == this.cap {
			key := this.d.deleteLast()
			delete(this.m, key)
			this.size--
		}
		this.size++
		this.d.addFront(key)
	}

}

type Node struct {
	key  int
	next *Node
	prev *Node
}

type DoubleLinkedList struct {
	head *Node
	tail *Node
}

func (d *DoubleLinkedList) delete(key int) {
	cur := d.head
	for cur != nil && cur.key != key {
		cur = cur.next
	}
	if cur == nil {
		return
	}
	if d.tail == d.head {
		d.tail = nil
		d.head = nil
		return
	}
	if cur == d.head {
		d.head = d.head.next
		d.head.prev = nil
	} else if cur == d.tail {
		d.tail = d.tail.prev
		cur.prev.next = nil
	} else {
		cur.prev.next = cur.next
		cur.next.prev = cur.prev
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

func (d *DoubleLinkedList) append(key int) {
	newNode := &Node{
		key: key,
	}
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

func (d *DoubleLinkedList) addFront(key int) {
	if d.head == nil {
		d.append(key)
	} else {
		newNode := &Node{
			key:  key,
			next: d.head,
		}

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
