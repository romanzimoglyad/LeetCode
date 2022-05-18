package main

import "fmt"

func main() {
	tn6 := &TreeNode{
		Val:   4,
		Left:  nil,
		Right: nil,
	}

	tn3 := &TreeNode{
		Val:  2,
		Left: tn6,
	}
	/*		tn4 := &TreeNode{
			Val:   3,
			Left:  nil,
			Right: nil,
		}*/
	tn2 := &TreeNode{
		Val:   2,
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
	// 4,2,5,1,3
	res := isSymmetric(tn1)
	fmt.Println(res)
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isSymmetric(root *TreeNode) bool {
	if root == nil {
		return true
	}

	return isSym(root.Left, root.Right)
}
func isSym(left, right *TreeNode) bool {
	if left == nil && right == nil {
		return true
	}
	if left == nil || right == nil {
		return false
	}
	return left.Val == right.Val && isSym(left.Left, right.Right) && isSym(left.Right, right.Left)
}
