package main

import "fmt"

type Node struct {
	val   int
	left  *Node
	right *Node
}

func main() {
	n1 := &Node{val: 1}
	n2 := &Node{val: 2}
	n3 := &Node{val: 3}
	n4 := &Node{val: 4}
	n2.left = n1
	n2.right = n3
	n3.right = n4
	n5 := &Node{
		val: 3,
	}
	n5.right = n4
	fmt.Println(isSub(n2, n5))
}

func isSub(main *Node, sub *Node) bool {
	if main == nil && sub == nil {
		return true
	}
	if main == nil {
		return false
	}

	return main.val == sub.val && isSub(main.left, sub.left) && isSub(main.right, sub.right) || isSub(main.left, sub) || isSub(main.right, sub)

}
