package main

import "fmt"

func main() {
	/*st4 := &Node{
		Val: 4,
	}
	*/
	/*
		st3 := &Node{
			Val: 3,
		}*/

	/*	st2 := &Node{
		Val: 2,
	}*/
	st1 := &Node{

		Val: -1,
	}
	st1.Random = st1
	//	st1.Next = st2
	//st2.Next = st3
	//	st2.Random = st1
	//st3.Random = st1
	//st4.Next = st1

	res := copyRandomList(st1)
	fmt.Println(res)
}

type Node struct {
	Val    int
	Next   *Node
	Random *Node
}

func copyRandomList(head *Node) *Node {
	cur := head
	dump := &Node{
		Next: head,
	}
	prev := dump
	m := make(map[*Node]*Node)
	for cur != nil {
		var newNode *Node
		if v, ok := m[cur]; ok {
			newNode = v
		} else {
			newNode = &Node{
				Val: cur.Val,
			}
		}
		m[cur] = newNode
		if cur.Random != nil {
			if r, ok := m[cur.Random]; ok {
				newNode.Random = r
			} else {
				newRandom := &Node{
					Val: cur.Random.Val,
				}
				newNode.Random = newRandom
				m[cur.Random] = newRandom
			}
		}

		prev.Next = newNode
		prev = newNode
		cur = cur.Next
	}
	return dump.Next
}
