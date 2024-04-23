package main

import (
	"fmt"
	"math"
)

func main() {
	g := [][]int{
		{0, 4, 0, 0, 0, 0, 0, 8, 0},
		{4, 0, 8, 0, 0, 0, 0, 11, 0},
		{0, 8, 0, 7, 0, 4, 0, 0, 2},
		{0, 0, 7, 0, 9, 14, 0, 0, 0},
		{0, 0, 0, 9, 0, 10, 0, 0, 0},
		{0, 0, 4, 14, 10, 0, 2, 0, 0},
		{0, 0, 0, 0, 0, 2, 0, 1, 6},
		{8, 11, 0, 0, 0, 0, 1, 0, 7},
		{0, 0, 2, 0, 0, 0, 6, 7, 0}}

	fmt.Println(d(g, 0))
}

func minDistance(dist []int, m map[int]bool) int {
	// Initialize min value
	min := math.MaxInt
	min_index := -1
	for v := 0; v < len(dist); v++ {
		if !m[v] && dist[v] <= min {
			min = dist[v]
			min_index = v
		}
	}

	return min_index
}
func d(graph [][]int, start int) []int {
	numberOfVertex := len(graph)
	dist := make([]int, numberOfVertex)
	dist[0] = 0
	for i := 1; i < numberOfVertex; i++ {
		dist[i] = math.MaxInt
	}
	m := make(map[int]bool)
	//q := NewPriority()
	//	q.push(start, start)
	for count := 0; count < numberOfVertex; count++ {
		//el := q.top()
		//if m[el.val] {
		//	continue
		//}
		u := minDistance(dist, m)
		m[u] = true
		for i := 0; i < numberOfVertex; i++ {
			if !m[i] && graph[u][i] != 0 && dist[u]+graph[u][i] < dist[i] {
				dist[i] = dist[u] + graph[u][i]

			}
		}

	}
	return dist
}

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

	if priority < p.head.priority {
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
