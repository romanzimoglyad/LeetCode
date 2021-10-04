package main

import "fmt"

func main() {
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
		Left:  nil,
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
		Right: tn3,
	}
	//1,2,4,5,3
	res := preorderTraversal(tn1)
	fmt.Println(res)
}

func preorderTraversal(root *TreeNode) []int {
	out := []int{}
	rec(&out, root)
	return out
}

func rec(out *[]int, root *TreeNode) {
	if root == nil {
		return
	}
	*out = append(*out, root.Val)

	if root.Left != nil {
		rec(out, root.Left)
	}

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
