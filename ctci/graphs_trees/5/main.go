package main

import (
	"fmt"
	"math"
)

type Node struct {
	val   int
	left  *Node
	right *Node
}

func main() {

	a100 := &Node{
		val:   100,
		left:  nil,
		right: nil,
	}
	a50 := &Node{
		val:   50,
		left:  nil,
		right: nil,
	}
	a150 := &Node{
		val:   100,
		left:  nil,
		right: nil,
	}
	a25 := &Node{
		val:   25,
		left:  nil,
		right: nil,
	}
	a15 := &Node{
		val:   15,
		left:  nil,
		right: nil,
	}

	a500 := &Node{
		val:   16,
		left:  nil,
		right: nil,
	}
	a100.left = a50
	a100.right = a150
	a50.left = a25
	a25.left = a15
	a15.right = a500
	fmt.Println(checkIfSearch(a100))
}

func checkIfSearch(in *Node) bool {
	if in == nil {
		return true
	}

	return checkIfSearchHelper(in, math.MinInt32, math.MaxInt32)
}

func checkIfSearchHelper(in *Node, min, max int) bool {
	if in == nil {
		return true
	}

	if in.right != nil && (in.right.val < in.val || in.right.val < min || in.right.val > max) {
		return false
	}

	if in.left != nil && (in.left.val > in.val || in.left.val > max || in.left.val < min) {
		return false
	}

	return checkIfSearchHelper(in.left, min, in.val) && checkIfSearchHelper(in.right, in.val, max)
}
