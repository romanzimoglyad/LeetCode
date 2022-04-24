package main

import "fmt"

func main() {
	n3 := Node{Val: 3}
	n4 := Node{Val: 4}
	n2 := Node{Val: 2}
	n1 := Node{Val: 1}
	n1.Neighbors = []*Node{&n2, &n4}
	n2.Neighbors = []*Node{&n1, &n3}
	n3.Neighbors = []*Node{&n2, &n4}
	n4.Neighbors = []*Node{&n3, &n1}

	g := cloneGraph(&n1)
	fmt.Println(g)
}

type Node struct {
	Val       int
	Neighbors []*Node
}

func cloneGraph(node *Node) *Node {

	visisted := make(map[int]*Node)

	return cl(visisted, node)
}

func cl(visisted map[int]*Node, node *Node) *Node {
	if node == nil {
		return nil
	}
	if n, ok := visisted[node.Val]; ok {
		return n
	}

	n := &Node{Val: node.Val}
	visisted[n.Val] = n
	for _, neighbor := range node.Neighbors {

		n.Neighbors = append(n.Neighbors, cl(visisted, neighbor))

	}
	return n
}
