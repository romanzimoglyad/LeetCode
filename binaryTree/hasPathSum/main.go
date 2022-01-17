package main

import "fmt"

var res int

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

	tn111 := &TreeNode{
		Val:   2,
		Left:  nil,
		Right: nil,
	}
	tn1 := &TreeNode{
		Val:   1,
		Left:  tn111,
		Right: nil,
	}
	//1 + 2  4 3
	res := hasPathSum(tn1, 1)
	fmt.Println(res)
}
func hasPathSum(root *TreeNode, targetSum int) bool {
	if root == nil {
		return false
	}
	if root.Val == targetSum && root.Left == nil && root.Right == nil {
		return true
	}
	if root.Left == nil && root.Right == nil {
		return false
	}

	return hasPathSum(root.Left, targetSum-root.Val) || hasPathSum(root.Right, targetSum-root.Val)
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
