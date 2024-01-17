package main

import "fmt"

func main() {
	n4 := &Node{
		val: 4,
	}
	n3 := &Node{
		val:  3,
		Next: n4,
	}
	n2 := &Node{
		val:  2,
		Next: n3,
	}
	n1 := &Node{
		val:  1,
		Next: n3,
	}
	n4.Next = n2

	fmt.Println(findLoop(n1))
}

type Node struct {
	val  int
	Next *Node
}

func findLoop(l1 *Node) *Node {
	slow := l1
	fast := l1
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
		if slow == fast {
			break
		}
	}

	if fast == nil || fast.Next == nil {
		return nil
	}
	fast = l1

	for slow != fast {
		slow = slow.Next
		fast = fast.Next
	}

	return slow
}

func reverse(n *Node) *Node {
	if n == nil || n.Next == nil {
		return n
	}
	end := reverse(n.Next)
	n.Next.Next = n
	n.Next = nil
	return end
}
