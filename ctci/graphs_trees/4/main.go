package main

func main() {

}

type Node struct {
	val   int
	left  *Node
	right *Node
}

func isBalanced(root *Node) bool {
	if root == nil {
		return true
	}
	if root.left == nil && root.right == nil {
		return true
	}
	return isBalanced(root.left) && isBalanced(root.right) && height(root.left)-height(root.right) < 2
}

func height(root *Node) int {
	if root == nil {
		return 0
	}

	return max(height(root.left), height(root.right)) + 1
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
