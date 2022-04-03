package main

import "fmt"

func main() {

	n2 := &TreeNode{
		Val:   2,
		Left:  nil,
		Right: nil,
	}
	n1 := &TreeNode{
		Val:  1,
		Left: n2,
	}
	res := lowestCommonAncestor(n1, n1, n2)
	fmt.Println(res.Val)
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var tmp *TreeNode

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	find(root, p, q)
	return tmp
}

func find(root, p, q *TreeNode) bool {
	if root == nil {
		return false
	}
	var left, right, mid int
	if find(root.Left, p, q) {
		left = 1
	}
	if find(root.Right, p, q) {
		right = 1
	}
	if root.Val == p.Val || root.Val == q.Val {
		mid = 1
	}
	if left+right+mid >= 2 {
		tmp = root
	}

	return left+right+mid > 0
}

func findOne(root, p *TreeNode, depth int) bool {
	if root == nil {
		return false
	}
	var found bool
	if root.Val == p.Val {
		found = true
	}
	depth++
	return found || findOne(root.Left, p, depth) || findOne(root.Right, p, depth)
}
