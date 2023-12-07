package main

import "fmt"

func main() {
	tn1 := &TreeNode{
		Val:   2,
		Left:  nil,
		Right: nil,
	}
	tn2 := &TreeNode{
		Val:   3,
		Left:  nil,
		Right: nil,
	}
	tn4 := &TreeNode{
		Val:   4,
		Left:  nil,
		Right: nil}
	root := &TreeNode{
		Val:   1,
		Left:  tn1,
		Right: tn2,
	}
	tn2.Right = tn4
	resp := inorderTraversal(root)
	fmt.Println(resp)
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func inorderTraversal(root *TreeNode) []int {
	resp := []int{}
	s := NewStack()
	cur := root
	for cur != nil || s.Len() > 0 {

		for cur != nil {
			s.Push(cur)
			cur = cur.Left
		}

		val := s.Pop()
		resp = append(resp, val.Val)
		cur = val.Right

	}
	return resp
}

func NewStack() stack {
	return stack{nodes: []*TreeNode{}}
}

type stack struct {
	nodes []*TreeNode
}

func (s *stack) Push(node *TreeNode) {

	s.nodes = append(s.nodes, node)
}

func (s *stack) Pop() *TreeNode {
	if len(s.nodes) == 0 {
		return nil
	}
	resp := s.nodes[len(s.nodes)-1]
	s.nodes = s.nodes[0 : len(s.nodes)-1]
	return resp
}

func (s *stack) Len() int { return len(s.nodes) }
