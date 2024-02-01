package main

import "fmt"

type Node struct {
	val    string
	childs []*Node
}

type Graph struct {
	nodes []*Node
}

func main() {

	a := &Node{
		val: "a",
	}
	b := &Node{
		val: "b",
	}
	c := &Node{
		val: "c",
	}
	d := &Node{
		val: "d",
	}
	a.childs = []*Node{b}
	b.childs = []*Node{c}
	c.childs = []*Node{d}
	fmt.Println(isPath2(a, d))
}

func isPath2(a, b *Node) bool {
	m1 := make(map[*Node]struct{})
	m2 := make(map[*Node]struct{})
	q1 := Queue{arr: make([]*Node, 0)}
	q2 := Queue{arr: make([]*Node, 0)}
	q1.push(a)
	q2.push(b)
	m1[a] = struct{}{}
	m2[b] = struct{}{}
	for q1.len() != 0 || q2.len() != 0 {
		el1 := q1.pop()
		if el1 != nil {
			for _, v := range el1.childs {
				if _, ok := m1[v]; !ok {
					if _, ok := m2[v]; ok {
						return true
					}

					m1[v] = struct{}{}
					q1.push(v)
				}
			}
		}
		el2 := q2.pop()
		if el2 != nil {
			for _, v := range el2.childs {
				if _, ok := m2[v]; !ok {
					if _, ok := m1[v]; ok {
						return true
					}

					m2[v] = struct{}{}
					q1.push(v)
				}
			}
		}
	}
	return false
}

func isPath(a, b *Node) bool {
	m := make(map[*Node]struct{})
	q := Queue{arr: make([]*Node, 0)}
	q.push(a)
	m[a] = struct{}{}
	for q.len() != 0 {
		el := q.pop()
		for _, v := range el.childs {
			if _, ok := m[v]; !ok {
				if v == b {
					return true
				}
				m[v] = struct{}{}
				q.push(v)
			}
		}
	}
	return false
}

type Queue struct {
	arr []*Node
}

func (q *Queue) push(n *Node) {
	q.arr = append(q.arr, n)
}

func (q *Queue) pop() *Node {
	if q.len() == 0 {
		return nil
	}
	el := q.arr[0]
	q.arr = q.arr[1:]
	return el
}

func (q *Queue) len() int {
	return len(q.arr)
}
