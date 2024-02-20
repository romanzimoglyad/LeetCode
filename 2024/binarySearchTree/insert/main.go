package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func insertIntoBST(root *TreeNode, val int) *TreeNode {
	if root == nil {
		return &TreeNode{Val: val}
	}
	cur := root
	for cur != nil {
		if cur.Val > val {
			if cur.Left == nil {
				cur.Left = &TreeNode{Val: val}
				return root
			}
			cur = cur.Left
		} else {
			if cur.Right == nil {
				cur.Right = &TreeNode{Val: val}
				return root
			}
			cur = cur.Right
		}
	}
	return root
}
