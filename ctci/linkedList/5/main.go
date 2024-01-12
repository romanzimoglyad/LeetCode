package main

import "fmt"

func main() {

	n3 := &Node{
		val: 1,
	}
	n2 := &Node{
		val:  2,
		Next: n3,
	}
	n1 := &Node{
		val:  3,
		Next: n2,
	}

	nn1 := &Node{
		val: 5,
	}
	res := addRevers(n1, nn1)
	fmt.Println(res)
}

type Node struct {
	val  int
	Next *Node
}

func add(v1, v2 *Node) *Node {
	rem := 0
	dum := &Node{}
	cur := dum
	for v1 != nil || v2 != nil {
		var res int
		if v2 == nil {
			res = v1.val + rem
			cur.Next = &Node{val: res % 10}
			v1 = v1.Next
		} else if v1 == nil {
			res = v1.val + +rem
			cur.Next = &Node{val: res % 10}
			v2 = v2.Next
		} else {
			res = v1.val + v2.val + rem
			cur.Next = &Node{val: res % 10}
			v1 = v1.Next
			v2 = v2.Next
		}
		cur = cur.Next
		rem = res / 10
	}

	if rem == 1 {
		cur.Next = &Node{val: 1}
	}

	return dum.Next
}

func addRevers(v1, v2 *Node) *Node {
	v1 = reverse(v1)
	v2 = reverse(v2)
	rem := 0
	dum := &Node{}
	cur := dum
	for v1 != nil || v2 != nil {
		var res int
		if v2 == nil {
			res = v1.val + rem
			cur.Next = &Node{val: res % 10}
			v1 = v1.Next
		} else if v1 == nil {
			res = v1.val + +rem
			cur.Next = &Node{val: res % 10}
			v2 = v2.Next
		} else {
			res = v1.val + v2.val + rem
			cur.Next = &Node{val: res % 10}
			v1 = v1.Next
			v2 = v2.Next
		}
		cur = cur.Next
		rem = res / 10
	}

	if rem == 1 {
		cur.Next = &Node{val: 1}
	}

	return dum.Next
}

func reverse(n *Node) *Node {
	if n == nil || n.Next == nil {
		return n
	}
	end := reverse(n.Next)
	n.Next.Next = n
	return end
}
