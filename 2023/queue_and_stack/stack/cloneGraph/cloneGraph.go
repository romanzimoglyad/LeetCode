package main

import "fmt"

func main() {

	n3 := &Node{
		Val: 3,
	}
	n2 := &Node{
		Val: 2,
	}
	n4 := &Node{
		Val: 4,
	}
	n1 := &Node{
		Val:       1,
		Neighbors: []*Node{n2, n4},
	}
	n2.Neighbors = []*Node{n1, n3}
	n3.Neighbors = []*Node{n2, n4}
	n4.Neighbors = []*Node{n1, n3}
	resp := cloneGraph(n1)
	fmt.Println(resp)
}

type Node struct {
	Val       int
	Neighbors []*Node
}

func cloneGraph(node *Node) *Node {
	if node == nil {
		return nil
	}

	s := Constructor()
	m := make(map[*Node]*Node)

	s.Push(node)
	for s.Len() > 0 {
		el := s.Top()
		s.Pop()
		var newNode *Node
		if newOrig, ok := m[el]; !ok {
			newNode = &Node{Val: el.Val}
			m[el] = newNode
		} else {
			newNode = newOrig
		}
		for _, v := range el.Neighbors {
			if newNeighnor, ok := m[v]; ok {
				newNode.Neighbors = append(newNode.Neighbors, newNeighnor)
			} else {
				newNeighbour := &Node{Val: v.Val}
				newNode.Neighbors = append(newNode.Neighbors, newNeighbour)
				m[v] = newNeighbour
				s.Push(v)
			}
		}

	}
	return m[node]
}

type stack struct {
	arr    []*Node
	length int
}

func Constructor() stack {
	return stack{arr: make([]*Node, 0)}
}

func (this *stack) Push(val *Node) {
	this.arr = append(this.arr, val)

	this.length++
}

func (this *stack) Pop() {

	this.arr = this.arr[:this.length-1]
	this.length--
}

func (this *stack) Top() *Node {
	return this.arr[this.length-1]
}

func (this *stack) Len() int {
	return this.length
}
