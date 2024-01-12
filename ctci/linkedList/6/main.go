package main

import "fmt"

func main() {
	n4 := &Node{
		val: 1,
	}
	n3 := &Node{
		val:  2,
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
	fmt.Println(isPalindtrom(n1))
}

type Node struct {
	val  int
	Next *Node
}

func isPalindtrom(n *Node) bool {
	length := 0
	cur := n
	for cur != nil {
		length++
		cur = cur.Next
	}
	midInd := 0
	if length%2 == 0 {
		midInd = length / 2
	} else {
		midInd = length/2 + 1
	}
	first := n
	second := n
	for i := 0; i < midInd; i++ {
		second = second.Next
	}
	newSec := reverse(second)
	for newSec != nil {
		if first.val != newSec.val {
			return false
		}
		first = first.Next
		newSec = newSec.Next
	}
	return true
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
