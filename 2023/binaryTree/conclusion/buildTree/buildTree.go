package main

import "fmt"

func main() {
	r7 := &TreeNode{
		Val: 7,
	}
	l6 := &TreeNode{
		Val: 6,
	}
	r5 := &TreeNode{
		Val: 5,
	}
	l4 := &TreeNode{
		Val: 4,
	}
	r3 := &TreeNode{
		Val:   3,
		Left:  l6,
		Right: r7,
	}
	l2 := &TreeNode{
		Val:   2,
		Left:  l4,
		Right: r5,
	}
	root1 := &TreeNode{
		Val:   1,
		Left:  l2,
		Right: r3,
	}
	fmt.Println(root1)
	r := buildTree([]int{4, 2, 5, 1, 6, 3, 7}, []int{4, 5, 2, 6, 7, 3, 1})
	fmt.Println(r)
	//r := buildTree([]int{1, 2}, []int{2, 1})
	//fmt.Println(r)
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func buildTree(inorder []int, postorder []int) *TreeNode {
	var resp *TreeNode
	inorderValToKey := make(map[int]int)
	for k, v := range inorder {
		inorderValToKey[v] = k
	}
	postorderInd := len(postorder) - 1
	//cur := postorder[postorderInd]
	//inorderStartKey := inorderValToKey[cur]
	// inorderStartKey 3
	// postorderInd 6
	resp = help(inorder, postorder, 0, len(postorder)-1, &postorderInd, inorderValToKey)
	/*	fmt.Println(rootKey)
		rightSubTreelen := len(inorder) - rootKey - 1
	*/

	return resp
}

// inorderStartKey 3 inorderLastKey 5
func help(inorder []int, postorder []int, inorderStartKey int, inorderLastKey int, postorderInd *int, inorderValToKey map[int]int) *TreeNode {

	if inorderStartKey > inorderLastKey {
		return nil
	}
	curInd := inorderValToKey[postorder[*postorderInd]]
	node := &TreeNode{
		Val: postorder[*postorderInd],
	}
	*postorderInd--

	node.Right = help(inorder, postorder, curInd+1, inorderLastKey, postorderInd, inorderValToKey)
	node.Left = help(inorder, postorder, inorderStartKey, curInd-1, postorderInd, inorderValToKey)

	return node
}

func help11(inorder []int, postorder []int, inorderStartKey int, inorderLastKey int, postorderInd *int, inorderValToKey map[int]int) *TreeNode {
	curInd := inorderValToKey[postorder[*postorderInd]]
	node := &TreeNode{
		Val: postorder[*postorderInd],
	}

	*postorderInd--
	if curInd == inorderLastKey {
		return node
	}

	if inorderStartKey == curInd {
		return node
	}
	node.Right = help(inorder, postorder, curInd+1, inorderLastKey, postorderInd, inorderValToKey)
	node.Left = help(inorder, postorder, inorderStartKey, curInd-1, postorderInd, inorderValToKey)
	return node
}

func help111(inorder []int, postorder []int, inorderStartKey int, inorderLastKey int, postorderInd *int, inorderValToKey map[int]int) *TreeNode {

	curInd := inorderValToKey[postorder[*postorderInd]]
	node := &TreeNode{
		Val: postorder[*postorderInd],
	}
	if inorderStartKey == inorderLastKey {
		return node
	}
	if inorderStartKey < 0 || inorderLastKey < 0 || inorderStartKey > inorderLastKey {
		return nil
	}
	*postorderInd--

	node.Right = help(inorder, postorder, curInd+1, inorderLastKey, postorderInd, inorderValToKey)
	node.Left = help(inorder, postorder, inorderStartKey, curInd-1, postorderInd, inorderValToKey)

	if curInd == inorderLastKey {
		return node
	}

	if inorderStartKey == curInd {
		return node
	}
	return node
}
