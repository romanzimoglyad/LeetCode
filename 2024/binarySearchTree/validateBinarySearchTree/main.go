package main

import (
	"fmt"
	"math"
)

func main() {
	n2 := &TreeNode{
		Val:   2,
		Left:  nil,
		Right: nil,
	}
	n3 := &TreeNode{
		Val:   3,
		Left:  nil,
		Right: nil,
	}

	n1 := &TreeNode{
		Val:   1,
		Left:  nil,
		Right: nil,
	}
	n2.Left = n1
	n2.Right = n3
	fmt.Println(isValidBST2(n2))

}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isValidBST(root *TreeNode) bool {
	arr := []int{}
	arr = inOrder(arr, root)

	for i := 0; i < len(arr)-1; i++ {
		if arr[i] >= arr[i+1] {
			return false
		}
	}
	return true
}

func inOrder(resp []int, root *TreeNode) []int {
	if root == nil {
		return resp
	}
	resp = inOrder(resp, root.Left)
	resp = append(resp, root.Val)
	resp = inOrder(resp, root.Right)
	return resp
}

func isValidBST2(root *TreeNode) bool {
	return check(root, math.MinInt, math.MaxInt)
}
func check(n *TreeNode, min, max int) bool {
	if n == nil {
		return true
	}
	if n.Val <= min || n.Val >= max {
		return false
	}
	ret1 := check(n.Left, min, n.Val)
	ret2 := check(n.Right, n.Val, max)

	return ret1 && ret2

}
