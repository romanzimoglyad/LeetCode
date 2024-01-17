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
		Next: n2,
	}
	n22 := &Node{
		val:  22,
		Next: nil,
	}
	n11 := &Node{
		val:  11,
		Next: n22,
	}
	fmt.Println(findInter(n1, n11))
}

type Node struct {
	val  int
	Next *Node
}

func findInter(l1, l2 *Node) *Node {
	cur1 := l1
	cur2 := l2
	for cur1 != cur2 {
		cur1 = cur1.Next
		if cur1 == nil {
			cur1 = l2
		}
		cur2 = cur2.Next
		if cur2 == nil {
			cur2 = l1
		}
		if cur1 == l2 && cur2 == l1 {
			return nil
		}
	}
	return cur1
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
