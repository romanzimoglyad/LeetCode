package main

import "fmt"

func main() {
	st2 := &Node{
		Val: 2,

		Next:  nil,
		Child: nil,
	}
	st1 := &Node{
		Val:   1,
		Prev:  nil,
		Next:  st2,
		Child: nil,
	}
	st2.Prev = st1
	st4 := &Node{
		Val: 4,
	}
	st3 := &Node{
		Val:  3,
		Prev: nil,
		//	Next:  st4,
		Child: nil,
	}
	st2.Child = st3
	st4.Prev = st3
	st5 := &Node{
		Val: 5,
	}
	st4.Child = st5
	st6 := &Node{
		Val:  6,
		Prev: st5,
	}
	st5.Next = st6
	//	st2.Next = st6
	st7 := &Node{
		Val:  7,
		Prev: st2,
	}
	st2.Next = st7
	res := flatten(st1)
	fmt.Println(res)
}

type Node struct {
	Val   int
	Prev  *Node
	Next  *Node
	Child *Node
}

func flatten(root *Node) *Node {
	if root == nil {
		return nil
	}
	dum := &Node{Next: &Node{}}
	handleDeep(dum, root)
	dum.Next.Prev = nil
	return dum.Next
}

// 1  2 33 333 222 11
func handleDeep(prev *Node, node *Node) *Node {
	if node == nil {
		return prev
	}
	node.Prev = prev
	prev.Next = node
	next := node.Next
	var tail *Node

	tail = handleDeep(node, node.Child)
	node.Child = nil

	return handleDeep(tail, next)

}

func flatten1(root *Node) *Node {
	resp := rr(root)
	tmp := resp
	for tmp != nil {
		tmp.Child = nil
		tmp = tmp.Next
	}
	return resp
}

func rr(root *Node) *Node {
	cur := root

	newNode := root
	first := newNode
	for cur != nil {
		var next *Node
		if cur.Next != nil {
			next = &Node{
				Val:   cur.Next.Val,
				Next:  cur.Next.Next,
				Child: cur.Next.Child,
			}
		}

		if cur.Child != nil {
			newNode.Next = rr(cur.Child)
			newNode.Next.Prev = newNode
			for newNode.Next != nil {
				newNode = newNode.Next
			}
		}
		if next != nil {
			newNode.Next = next
			newNode.Next.Prev = newNode
			newNode = newNode.Next
		}
		cur = next
	}
	return first
}
