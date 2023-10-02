package main

import "fmt"

func main() {

	node1 := &TreeNode{
		Val:   2,
		Left:  nil,
		Right: nil,
	}
	node2 := &TreeNode{
		Val:   4,
		Left:  nil,
		Right: nil,
	}
	node3 := &TreeNode{
		Val:   3,
		Left:  nil,
		Right: nil,
	}
	node4 := &TreeNode{
		Val:   2,
		Left:  nil,
		Right: nil,
	}
	node5 := &TreeNode{
		Val:   4,
		Left:  nil,
		Right: nil,
	}
	node6 := &TreeNode{
		Val:   4,
		Left:  nil,
		Right: nil,
	}
	root := &TreeNode{
		Val:   1,
		Left:  node1,
		Right: node3,
	}
	node1.Left = node2
	node3.Left = node4
	node3.Right = node6
	node4.Left = node5
	resp := findDuplicateSubtrees(root)
	fmt.Println(resp)

}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var m map[string]int
var response []*TreeNode

func findDuplicateSubtrees(root *TreeNode) []*TreeNode {
	if root == nil {
		return nil
	}
	m = make(map[string]int)
	response = []*TreeNode{}
	getSubTree1(root)

	return response
}

func getSubTree1(root *TreeNode) string {
	if root == nil {
		return ""
	}
	resp := "(" + getSubTree1(root.Left) + ")" + nodeToString(root) + "(" +
		getSubTree1(root.Right) + ")"
	m[resp]++
	if m[resp] == 2 {
		response = append(response, root)
	}
	return resp
}

func nodeToString(node *TreeNode) string {
	return string(node.Val)
}
