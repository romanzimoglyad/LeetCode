package main

import "fmt"

func main() {

	n4 := &Node{
		val: 4,
	}
	n3 := &Node{
		val:  1,
		Next: n4,
	}
	n2 := &Node{
		val:  1,
		Next: n3,
	}
	n1 := &Node{
		val:  1,
		Next: n2,
	}
	root := &Node{
		val:  1,
		Next: n1,
	}
	removeDupsV2(root)
	fmt.Println(root)
}

type Node struct {
	val  int
	Next *Node
}

func removeDupsV2(root *Node) {
	slowCur := root

	for slowCur != nil && slowCur.Next != nil {
		fastCur := slowCur
		for fastCur != nil && fastCur.Next != nil {
			if slowCur.val == fastCur.Next.val {
				fastCur.Next = fastCur.Next.Next
			} else {
				fastCur = fastCur.Next
			}
		}
		slowCur = slowCur.Next
	}
}
func removeDups(root *Node) {
	m := make(map[int]struct{})
	cur := root
	m[cur.val] = struct{}{}
	for cur != nil && cur.Next != nil {
		if _, ok := m[cur.Next.val]; ok {
			cur.Next = cur.Next.Next
		} else {
			m[cur.Next.val] = struct{}{}
		}
		cur = cur.Next
	}

}
