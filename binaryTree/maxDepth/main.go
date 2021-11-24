package main

import "fmt"

var res int

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

	tn3 := &TreeNode{
		Val:   3,
		Left:  tn6,
		Right: tn7,
	}
	tn4 := &TreeNode{
		Val:   4,
		Left:  tn3,
		Right: nil,
	}
	tn2 := &TreeNode{
		Val:   2,
		Left:  tn4,
		Right: tn5,
	}
	tn1 := &TreeNode{
		Val:   1,
		Left:  tn2,
		Right: nil,
	}
	//4,2,5,1,3
	res := maxDepth(tn1)
	fmt.Println(res)
}
func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}

	left := maxDepth(root.Left)
	right := maxDepth(root.Right)
	return Max(left, right) + 1
}
func Max(x, y int) int {
	if x < y {
		return y
	}
	return x
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
