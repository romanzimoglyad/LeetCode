package main

import "fmt"

func main() {
	tn5 := &Node{
		Val:   5,
		Left:  nil,
		Right: nil,
	}
	tn6 := &Node{
		Val:   6,
		Left:  nil,
		Right: nil,
	}
	tn7 := &Node{
		Val:   7,
		Left:  nil,
		Right: nil,
	}
	tn4 := &Node{
		Val:   4,
		Left:  nil,
		Right: nil,
	}
	tn3 := &Node{
		Val:   3,
		Left:  tn6,
		Right: tn7}

	tn2 := &Node{
		Val:   2,
		Left:  tn4,
		Right: tn5,
	}
	tn1 := &Node{
		Val:   1,
		Left:  tn2,
		Right: tn3,
	}

	resp := connect(tn1)
	fmt.Println(resp)
}

type Node struct {
	Val   int
	Left  *Node
	Right *Node
	Next  *Node
}

func connect(root *Node) *Node {
	if root == nil {
		return nil
	}
	q := queue{}
	q.Push(root)
	for q.Len() != 0 {
		l := q.Len()
		tmp := &Node{}
		for i := 0; i < l; i++ {

			node := q.Pop()
			if i != 0 {
				tmp.Next = node
			}
			tmp = node
			fmt.Print(node.Val)
			if node.Left != nil {
				q.Push(node.Left)
			}
			if node.Right != nil {
				q.Push(node.Right)
			}
		}
		fmt.Println("")
	}
	return root
}

type queue struct {
	arr []*Node
}

func (q *queue) Len() int { return len(q.arr) }
func (q *queue) Push(el *Node) {
	q.arr = append(q.arr, el)
}
func (q *queue) Pop() *Node {
	if q.Len() == 0 {
		return nil
	}
	el := q.arr[0]
	q.arr = q.arr[1:]
	return el
}
