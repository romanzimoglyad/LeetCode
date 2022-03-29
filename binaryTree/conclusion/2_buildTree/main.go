package main

import "fmt"

func main() {

	res := buildTree([]int{1, 2, 4, 5, 3, 6, 7}, []int{4, 2, 5, 1, 6, 3, 7})
	fmt.Println(res)
}
func buildTree(preorder []int, inorder []int) *TreeNode {
	inOrderMap := make(map[int]int)
	for key, value := range inorder {
		inOrderMap[value] = key
	}
	var cur int
	root := build(inOrderMap, &cur, 0, len(inorder)-1, preorder)

	return root

}

func build(inOrderMap map[int]int, cursor *int, start, end int, preOrder []int) *TreeNode {

	if start > end {
		return nil
	}
	node := &TreeNode{Val: preOrder[*cursor]}
	*cursor++
	ind := inOrderMap[node.Val]

	node.Left = build(inOrderMap, cursor, start, ind-1, preOrder)
	node.Right = build(inOrderMap, cursor, ind+1, end, preOrder)
	return node
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
