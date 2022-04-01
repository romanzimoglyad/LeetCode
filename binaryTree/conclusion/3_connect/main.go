package main

import "fmt"

func main() {
	n4 := &Node{
		Val: 4,
	}
	n5 := &Node{
		Val: 5,
	}
	n6 := &Node{
		Val: 6,
	}
	n7 := &Node{
		Val: 7,
	}
	n2 := &Node{
		Val:   2,
		Left:  n4,
		Right: n5,
	}
	n3 := &Node{
		Val:   3,
		Left:  n6,
		Right: n7,
	}
	n1 := &Node{
		Val:   1,
		Left:  n2,
		Right: n3,
	}
	res := connect(n1)
	fmt.Println(res)
}

type Node struct {
	Val   int
	Left  *Node
	Right *Node
	Next  *Node
}

func connect(root *Node) *Node {

	if root == nil {
		return root
	}
	q := newQueue()
	q.add([]*Node{root})
	for elem := q.pop(); elem != nil; elem = q.pop() {

		newLine := []*Node{}
		for i := range elem {
			if elem[i].Left != nil {
				newLine = append(newLine, elem[i].Left)
			}
			if elem[i].Right != nil {
				newLine = append(newLine, elem[i].Right)
			}
		}
		if len(newLine) > 0 {
			q.add(newLine)
		}
		if len(newLine) > 1 {
			for i := 0; i < len(newLine)-1; i++ {
				newLine[i].Next = newLine[i+1]
			}
		}

	}
	return root
}

type queue struct {
	nodes  [][]*Node
	length int
}

func newQueue() queue {
	return queue{
		nodes:  [][]*Node{},
		length: 0,
	}
}
func (q *queue) add(elem []*Node) {
	q.nodes = append(q.nodes, elem)
	q.length++
}
func (q *queue) pop() []*Node {
	if q.length == 0 {
		return nil
	}
	res := q.nodes[len(q.nodes)-1]
	q.nodes = q.nodes[0 : len(q.nodes)-1]
	q.length--
	return res
}
