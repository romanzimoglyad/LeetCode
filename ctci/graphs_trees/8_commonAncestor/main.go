package __commonAncestor

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {

}

func commonAncestor(root, n1, n2 *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	if root == n1 || root == n2 {
		return root
	}
	left := commonAncestor(root.Left, n1, n2)
	right := commonAncestor(root.Right, n1, n2)
	if left != nil && right != nil {
		return root
	}
	if left != nil {
		return left
	}
	return right

}
