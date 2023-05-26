package main

import "fmt"

func main() {
	/*st4 := &Node{
		Val: 2,
	}*/

	st1 := &Node{
		Val: 3,
	}
	st1.Next = st1
	res := insert(st1, 2)
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

	prev := aNode
	cur := aNode.Next
	for {
		if prev.Val > cur.Val {

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
	}
	return aNode
}
