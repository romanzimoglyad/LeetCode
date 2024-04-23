package main

import "fmt"

func main() {
	q := NewPriority()
	q.push(7, 7)
	q.push(6, 6)
	q.push(4, 4)
	q.push(5, 5)
	q.push(2, 2)
	fmt.Println(q.top())
	fmt.Println(q.top())
	fmt.Println(q.top())
	fmt.Println(q.top())
	fmt.Println(q.top())
	fmt.Println(q.top())
}

//4 2 7

type Node struct {
	val      int
	priority int
	Next     *Node
}
type Priority struct {
	head *Node
}

func NewPriority() *Priority {
	return &Priority{}
}

// 7 6  4
func (p *Priority) push(val, priority int) {
	if p.head == nil {
		p.head = &Node{
			val:      val,
			priority: priority,
		}
		return
	}
	cur := p.head
	prev := &Node{Next: cur}

	if priority <= p.head.priority {
		newNode := &Node{
			val:      val,
			priority: priority,
			Next:     p.head,
		}
		p.head = newNode
		return
	}

	for cur != nil && priority > cur.priority {
		prev = cur
		cur = cur.Next
	}
	tmpNext := prev.Next
	newNode := &Node{
		val:      val,
		priority: priority,
		Next:     tmpNext,
	}
	prev.Next = newNode
	return
}

func (p *Priority) top() *Node {
	if p.head == nil {
		return nil
	}
	next := p.head.Next
	head := p.head
	p.head = next
	return head

}

func (p *Priority) peek() *Node {
	return p.head
}
