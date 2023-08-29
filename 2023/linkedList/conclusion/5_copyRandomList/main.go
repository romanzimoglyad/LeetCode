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

	res := copyRandomList(st1)
	fmt.Println(res)
}

type Node struct {
	Val  int
	Next *Node
}

func copyRandomList(head *Node) *Node {

}
