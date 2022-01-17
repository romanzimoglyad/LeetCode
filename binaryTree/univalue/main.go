package main

import "fmt"

func main() {

	/*	tn7 := &TreeNode{
			Val:   3,
			Left:  nil,
			Right: nil,
		}
		tn6 := &TreeNode{
			Val:   4,
			Left:  tn7,
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
			}
		tn2 := &TreeNode{
			Val:   2,
			Left:  tn7,
			Right: tn6,
		}*/
	tn2 := &TreeNode{
		Val:   5,
		Left:  nil,
		Right: nil,
	}
	tn3 := &TreeNode{
		Val:   5,
		Left:  nil,
		Right: nil,
	}
	tn11 := &TreeNode{
		Val:   5,
		Left:  tn2,
		Right: tn3,
	}
	tn7 := &TreeNode{
		Val:   5,
		Left:  nil,
		Right: nil,
	}
	tn6 := &TreeNode{
		Val:   5,
		Left:  nil,
		Right: tn7,
	}
	tn1 := &TreeNode{
		Val:   5,
		Left:  tn11,
		Right: tn6,
	}
	//1 + 2  4 3

	res := countUnivalSubtrees(tn1)
	fmt.Println(res)
}
func countUnivalSubtrees(root *TreeNode) int {
	if root == nil {
		return 0
	}
	res := 0
	is(root, &res)
	return res
}
func is(root *TreeNode, res *int) bool {

	if root.Right == nil && root.Left == nil {
		*res++
		return true
	}
	isUnival := true
	if root.Left != nil {
		isUnival = is(root.Left, res) && isUnival && root.Left.Val == root.Val
	}
	if root.Right != nil {
		isUnival = is(root.Right, res) && isUnival && root.Right.Val == root.Val
	}

	if !isUnival {
		return false
	}
	*res++
	return true

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
