package main

import "fmt"

func main() {
	tn6 := &TreeNode{
		Val:   6,
		Left:  nil,
		Right: nil,
	}
	tn7 := &TreeNode{
		Val:   7,
		Left:  nil,
		Right: nil,
	}
	tn5 := &TreeNode{
		Val:   5,
		Left:  nil,
		Right: nil,
	}
	tn4 := &TreeNode{
		Val:   4,
		Left:  nil,
		Right: nil,
	}
	tn3 := &TreeNode{
		Val:   3,
		Left:  tn6,
		Right: tn7,
	}
	tn2 := &TreeNode{
		Val:   2,
		Left:  tn4,
		Right: tn5,
	}
	tn1 := &TreeNode{
		Val:   1,
		Left:  tn2,
		Right: tn3,
	}
	//1,2,3,4,5,6,7
	res := levelOrder(tn1)
	fmt.Println(res)
}

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
type queue struct {
	nodes  [][]*TreeNode
	length int
}

func newQueue() queue {
	return queue{
		nodes:  [][]*TreeNode{},
		length: 0,
	}
}
func (q *queue) add(elem []*TreeNode) {
	q.nodes = append(q.nodes, elem)
	q.length++
}
func (q *queue) pop() []*TreeNode {
	if q.length == 0 {
		return nil
	}
	res := q.nodes[len(q.nodes)-1]
	q.nodes = q.nodes[0 : len(q.nodes)-1]
	q.length--
	return res
}

func levelOrder(root *TreeNode) [][]int {
	res := [][]int{}
	if root == nil {
		return [][]int{}
	}
	q := newQueue()
	q.add([]*TreeNode{root})
	for elem := q.pop(); elem != nil; elem = q.pop() {
		res = append(res, convSlice(elem))
		newLine := []*TreeNode{}
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

	}
	return res
}

func convSlice(in []*TreeNode) []int {
	out := []int{}
	for i := range in {
		out = append(out, in[i].Val)
	}
	return out
}
func rec(out *[]int, root *TreeNode) {
	if root == nil {
		return
	}

	if root.Left != nil {
		rec(out, root.Left)
	}
	*out = append(*out, root.Val)
	if root.Right != nil {
		rec(out, root.Right)
	}

	return
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
