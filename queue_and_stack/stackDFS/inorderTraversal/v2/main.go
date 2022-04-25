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
	// 4,2,5,1,3
	res := inorderTraversal(tn1)
	fmt.Println(res)
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func inorderTraversal(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}
	var result []int
	stack := make([]*TreeNode, 0)

	current := root
	for current != nil || len(stack) != 0 {
		for current != nil {
			stack = append(stack, current)
			current = current.Left

		}
		current = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		result = append(result, current.Val)
		current = current.Right

	}
	return result
}
