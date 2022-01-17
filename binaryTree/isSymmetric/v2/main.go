package main

import "fmt"

var res int

func main() {
	tn6 := &TreeNode{
		Val:   4,
		Left:  nil,
		Right: nil,
	}
	tn7 := &TreeNode{
		Val:   3,
		Left:  nil,
		Right: nil,
	}
	tn5 := &TreeNode{
		Val:   3,
		Left:  nil,
		Right: nil,
	}

	tn3 := &TreeNode{
		Val:   2,
		Left:  tn6,
		Right: tn5,
	}
	/*		tn4 := &TreeNode{
			Val:   3,
			Left:  nil,
			Right: nil,
		}*/
	tn2 := &TreeNode{
		Val:   2,
		Left:  tn7,
		Right: tn6,
	}
	tn1 := &TreeNode{
		Val:   1,
		Left:  tn2,
		Right: tn3,
	}

	/*	tn22 := &TreeNode{
			Val:   22,
			Left:  nil,
			Right: nil,
		}
		tn33 := &TreeNode{
			Val:   22,
			Left:  nil,
			Right: nil,
		}
		tn11 := &TreeNode{
			Val:   11,
			Left:  tn22,
			Right: tn33,
		}
	*/
	//4,2,5,1,3
	res := isSymmetric(tn1)
	fmt.Println(res)
}
func isSymmetric(root *TreeNode) bool {
	if root == nil {
		return false
	}
	return sim(root.Left, root.Right)

}

func sim(left *TreeNode, right *TreeNode) bool {
	if left == nil && right == nil {
		return true
	}
	if left == nil || right == nil {
		return false
	}
	return left.Val == right.Val && sim(left.Left, right.Right) && sim(left.Right, right.Left)
}

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

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
