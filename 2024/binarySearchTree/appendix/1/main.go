package main

import (
	"fmt"
	"math"
)

func main() {
	//n3 := &TreeNode{Val: 3}
	//n9 := &TreeNode{Val: 9}
	//n20 := &TreeNode{Val: 20}
	//n15 := &TreeNode{Val: 15}
	//n7 := &TreeNode{Val: 7_build_prder}
	//n3.Left = n9
	//n3.Right = n20
	//
	//n20.Left = n15
	//n20.Right = n7
	//fmt.Println(isBalanced(n3))
	//
	//
	n1 := &TreeNode{Val: 1}
	n2 := &TreeNode{Val: 2}
	n3 := &TreeNode{Val: 3}
	n4 := &TreeNode{Val: 4}
	n5 := &TreeNode{Val: 5}

	n8 := &TreeNode{Val: 8}
	n1.Left = n2
	n1.Right = n3
	n2.Left = n4
	n2.Right = n5
	n4.Left = n8

	fmt.Println(isBalanced(n1))
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isBalanced(root *TreeNode) bool {
	if root == nil {
		return true
	}
	return isBalanced(root.Right) && isBalanced(root.Left) && math.Abs(float64(travers(root.Left))-float64(travers(root.Right))) < 2
}

func travers(root *TreeNode) int {
	if root == nil {
		return -1
	}
	leftDepth := travers(root.Left)
	rightDepth := travers(root.Right)
	return max(leftDepth, rightDepth) + 1
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
