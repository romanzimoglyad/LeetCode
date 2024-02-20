package main

import "fmt"

func main() {
	n5 := &TreeNode{Val: 5}
	n3 := &TreeNode{Val: 3}
	n6 := &TreeNode{Val: 6}
	n2 := &TreeNode{Val: 2}
	n4 := &TreeNode{Val: 4}
	n7 := &TreeNode{Val: 7}

	n5.Left = n3
	n5.Right = n6
	n3.Left = n2
	n3.Right = n4
	n6.Right = n7

	re := deleteNode(n5, 3)
	fmt.Println(re)

}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func deleteNode(root *TreeNode, key int) *TreeNode {

	if root == nil {
		return root
	}
	if root.Val == key {
		if root.Left == nil && root.Right == nil {
			return nil
		}
		if root.Right != nil {
			s := successor(root)
			root.Val = s.Val
			root.Right = deleteNode(root.Right, s.Val)
		} else {
			s := predeccessor(root)
			root.Val = s.Val
			root.Left = deleteNode(root.Left, s.Val)
		}
	}
	if root.Val > key {
		root.Left = deleteNode(root.Left, key)
	} else if root.Val < key {
		root.Right = deleteNode(root.Right, key)
	}

	return root
}

func successor(node *TreeNode) *TreeNode {
	node = node.Right
	for node.Left != nil {
		node = node.Left
	}
	return node
}
func predeccessor(node *TreeNode) *TreeNode {
	node = node.Left
	for node.Right != nil {
		node = node.Right
	}
	return node
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
