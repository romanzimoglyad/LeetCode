package main

import (
	"errors"
	"fmt"
)

func main() {
	pr := []int{1, 2, 3, 4, 5, 6}
	deps := [][2]int{{1, 4}, {6, 2}, {2, 4}, {6, 1}, {4, 3}}
	resp, err := findDeps(pr, deps)
	fmt.Println(resp)
	fmt.Println(err)
}

type Node struct {
	val    int
	deps   []*Node
	parent []*Node
}

func findDeps(pr []int, deps [][2]int) ([]int, error) {
	m := make(map[int]*Node)
	for _, dep := range deps {
		var n *Node
		if node, ok := m[dep[0]]; ok {
			n = node
		} else {
			n = &Node{
				val: dep[0],
			}
			m[dep[0]] = n
		}
		var depToAdd *Node
		if existedDep, ok := m[dep[1]]; ok {
			depToAdd = existedDep
			depToAdd.parent = append(depToAdd.parent, n)
		} else {
			n := &Node{
				val:    dep[1],
				parent: []*Node{n},
			}
			m[dep[1]] = n
			depToAdd = n
		}
		n.deps = append(n.deps, depToAdd)
	}
	var resp []int
	var q []*Node
	seen := make(map[int]struct{})
	for _, v := range m {
		if v.parent == nil {
			q = append(q, v)
			for len(q) != 0 {
				el := q[0]
				q = q[1:]
				if _, ok := seen[el.val]; !ok {
					seen[el.val] = struct{}{}
					resp = append(resp, el.val)
					for _, dep := range el.deps {
						q = append(q, dep)
					}
				}
			}
		}
	}
	if len(resp) != len(pr) {
		return []int{}, errors.New("rtt")
	}

	return resp, nil
}
