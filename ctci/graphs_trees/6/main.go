package main

func main() {

}

type Node struct {
	val   int
	left  *Node
	right *Node
}

func inOrderSuccessor(in *Node, val int) *Node {
	var prev *Node
	cur := in
	for cur != nil {
		if cur.val > val {
			prev = cur
			cur = cur.left
		} else {
			cur = cur.right
		}
	}
	return prev
}
