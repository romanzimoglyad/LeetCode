package main

import "fmt"

func main() {
	/*st4 := &Node{
		Val: 4,
	}
	*/

	st3 := &Node{
		Val: 3,
	}

	st2 := &Node{
		Val: 3,
	}
	st1 := &Node{
		Val: 3,
	}
	st1.Next = st2
	st2.Next = st3
	st3.Next = st1

	//st4.Next = st1

	res := insert(st1, 0)
	fmt.Println(res)
}

type Node struct {
	Val  int
	Next *Node
}

func insert(aNode *Node, x int) *Node {
	if aNode == nil {
		n := &Node{
			Val: x,
		}
		n.Next = n
		return n
	}
	if aNode.Next == aNode {
		n := &Node{
			Val: x,
		}
		aNode.Next = n
		n.Next = aNode
		return aNode
	}
	startFirst := aNode
	prev := aNode
	cur := aNode.Next
	var start *Node
	for {
		if cur.Val < prev.Val || cur == startFirst {
			start = cur
			break
		}

		prev = prev.Next
		cur = cur.Next
	}

	for {
		if cur.Val >= x {
			newNode := &Node{
				Val: x,
			}

			newNode.Next = cur
			prev.Next = newNode
			break
		}
		prev = prev.Next
		cur = cur.Next
		if cur == start {
			newNode := &Node{
				Val: x,
			}
			newNode.Next = cur
			prev.Next = newNode
			break
		}
	}
	return aNode
}
