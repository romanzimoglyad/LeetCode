package main

import "fmt"

var res int

func main() {
	/*tn6 := &TreeNode{
		Val:   4,
		Left:  nil,
		Right: nil,
	}*/
	/*	tn7 := &TreeNode{
			Val:   3,
			Left:  nil,
			Right: nil,
		}
		tn5 := &TreeNode{
			Val:   3,
			Left:  nil,
			Right: nil,
		}
	*/
	/*tn3 := &TreeNode{
		Val:   2,
		Left:  nil,
		Right: nil,
	}
	tn4 := &TreeNode{
		Val:   3,
		Left:  nil,
		Right: nil,
	}*/
	/*	tn2 := &TreeNode{
		Val:   0,
		Left:  nil,
		Right: nil,
	}*/
	tn1 := &TreeNode{
		Val:   1,
		Left:  nil,
		Right: nil,
	}
	/*	fmt.Println(tn1)
		tn22 := &TreeNode{
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
		}*/

	//4,2,5,1,3
	res := isSymmetric(tn1)
	fmt.Println(res)
}
func isSymmetric(root *TreeNode) bool {
	if root == nil {
		return false
	}
	if root.Left == nil && root.Right == nil {
		return true
	}
	left := []*TreeNode{root.Left}
	if root.Left == nil {
		left = nil
	}
	right := []*TreeNode{root.Right}
	if root.Right == nil {
		right = nil
	}
	res := compareLayer(left, right)

	return res

}
func compareLayer(left []*TreeNode, right []*TreeNode) bool {
	if left == nil || right == nil {
		return false
	}
	if len(left) != len(right) {
		return false
	}
	for i := 0; i < len(left); i++ {
		if left[i].Val != right[i].Val {
			return false
		}
	}

	fake := &TreeNode{
		Val:   -1000,
		Left:  nil,
		Right: nil,
	}
	newLeft := []*TreeNode{}
	for _, elem := range left {
		if elem.Left != nil {
			newLeft = append(newLeft, elem.Left)
		} else {
			newLeft = append(newLeft, fake)
		}
		if elem.Right != nil {
			newLeft = append(newLeft, elem.Right)
		} else {
			newLeft = append(newLeft, fake)
		}
	}
	newRight := []*TreeNode{}
	for _, elem := range right {
		if elem.Right != nil {
			newRight = append(newRight, elem.Right)
		} else {
			newRight = append(newRight, fake)
		}
		if elem.Left != nil {
			newRight = append(newRight, elem.Left)
		} else {
			newRight = append(newRight, fake)
		}
	}
	if allFakes(newLeft) && allFakes(newRight) {
		return true
	}
	res := compareLayer(newLeft, newRight)
	return res
}
func left(root *TreeNode, leftArr *[]int) {
	if root == nil {
		return
	}
	left(root.Left, leftArr)

}

func allFakes(a []*TreeNode) bool {
	for i := range a {
		if a[i].Val != -1000 {
			return false
		}

	}
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
