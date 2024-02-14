package main

func main() {

}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func searchDST(root *TreeNode, val int) *TreeNode {
	if root == nil {
		return nil
	}
	if root.Val == val {
		return root
	}
	if root.Left == nil && root.Right == nil {
		return nil
	}
	if root.Val > val {
		return searchDST(root.Left, val)
	}
	return searchDST(root.Right, val)

}
