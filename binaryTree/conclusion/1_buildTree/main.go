package main

import "fmt"

func main() {

	res := buildTree([]int{9, 3, 15, 20, 7}, []int{9, 15, 7, 20, 3})
	fmt.Println(res)
}
func buildTree(inorder []int, postorder []int) *TreeNode {
	if len(inorder) == 1 {
		return &TreeNode{Val: inorder[0]}
	}
	revertPostOrder := revert(postorder)
	rootElem := &TreeNode{Val: revertPostOrder[0]}
	revertPostOrder = revertPostOrder[1:]
	var rootKey int
	for key, value := range inorder {
		if value == rootElem.Val {
			rootKey = key
		}
	}
	var leftArray []int
	var rightArray []int
	if rootKey != 0 {
		leftArray = inorder[:rootKey]
	}
	if rootKey != len(inorder)-1 {
		rightArray = inorder[rootKey+1:]
	}
	if len(rightArray) > 0 {
		postorderRightArray := revertPostOrder[:len(rightArray)]
		node := buildTree(rightArray, revert(postorderRightArray))
		rootElem.Right = node
	}
	if len(leftArray) > 0 {
		postorderLeftArray := revertPostOrder[len(rightArray):]
		node := buildTree(leftArray, revert(postorderLeftArray))
		rootElem.Left = node
	}
	return rootElem
}

func revert(in []int) []int {
	revertPostOrder := []int{}
	for i := len(in) - 1; i >= 0; i-- {
		revertPostOrder = append(revertPostOrder, in[i])
	}
	return revertPostOrder
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
