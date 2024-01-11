package main

import "fmt"

type Node struct {
	val  int
	Next *Node
}

func main() {
	n4 := &Node{
		val: 4,
	}
	n3 := &Node{
		val:  2,
		Next: n4,
	}
	n2 := &Node{
		val:  3,
		Next: n3,
	}
	n1 := &Node{
		val:  5,
		Next: n2,
	}
	root := &Node{
		val:  10,
		Next: n1,
	}
	res := partion(root, 4)
	fmt.Println(res)
}

func partion(cur *Node, val int) *Node {

	left := &Node{}
	start := left
	right := &Node{}
	rightStart := right

	for cur != nil {
		if cur.val < val {
			left.Next = &Node{
				val: cur.val,
			}
			left = left.Next
		} else {
			right.Next = &Node{
				val: cur.val,
			}
			right = right.Next
		}
		cur = cur.Next
	}
	left.Next = rightStart.Next
	return start.Next
}
