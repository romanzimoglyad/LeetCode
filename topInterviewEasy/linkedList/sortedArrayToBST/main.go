package main

import "fmt"

func main() {
	tn33 := &TreeNode{
		Val: 3,
	}
	tn333 := &TreeNode{
		Val: 7,
	}
	tn3 := &TreeNode{
		Val:   6,
		Left:  tn33,
		Right: tn333,
	}

	tn2 := &TreeNode{
		Val: 4,
	}
	tn1 := &TreeNode{
		Val:   5,
		Left:  tn2,
		Right: tn3,
	}
	// 1,2,3,4,5,6,7_build_prder
	res := isValidBST(tn1)
	fmt.Println(res)
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func sortedArrayToBST(nums []int) *TreeNode {

}
