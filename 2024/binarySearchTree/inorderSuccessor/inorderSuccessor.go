package main

import "fmt"

func main() {

	n1 := &TreeNode{Val: 1}
	n2 := &TreeNode{Val: 2}
	n3 := &TreeNode{Val: 3}
	n2.Left = n1
	n2.Right = n3
	resp := inorderSuccessor(n2, n1)
	fmt.Println(resp)

}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func inorderSuccessor(root *TreeNode, p *TreeNode) *TreeNode {

	cur := root
	var prev *TreeNode
	for cur != nil {
		if p.Val >= cur.Val {
			cur = cur.Right
		} else {
			prev = cur
			cur = cur.Left
		}
	}

	return prev
}

func inorderSuccessorV2(root *TreeNode, p *TreeNode) *TreeNode {
	arr := make([]*TreeNode, 0)
	bst(&arr, root)
	for k, v := range arr {
		if v.Val == p.Val {
			if k == len(arr)-1 {
				return nil
			}
			return arr[k+1]
		}
	}
	return nil
}

func bst(arr *[]*TreeNode, root *TreeNode) {
	if root == nil {
		return
	}

	bst(arr, root.Left)
	*arr = append(*arr, root)
	bst(arr, root.Right)
}
