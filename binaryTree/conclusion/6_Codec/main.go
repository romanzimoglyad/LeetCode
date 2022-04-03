package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	/*
		n4 := &TreeNode{
			Val:   4,
			Left:  nil,
			Right: nil,
		}
		n5 := &TreeNode{
			Val:   5,
			Left:  nil,
			Right: nil,
		}
		n3 := &TreeNode{
			Val:   3,
			Left:  n4,
			Right: n5,
		}
	*/

	n2 := &TreeNode{
		Val:   2,
		Left:  nil,
		Right: nil,
	}

	n1 := &TreeNode{
		Val:  1,
		Left: n2,
	}
	ser := Constructor()

	data := ser.serialize(n1)
	fmt.Println(data)
	ans := ser.deserialize(data)
	fmt.Println(ans)
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func cutLastNils(in string) string {
	for strings.HasSuffix(in, ",nil") {
		in = in[0 : len(in)-len(",nil")]
	}
	return in
}

type Codec struct {
	str       string
	root      *TreeNode
	lastLayer []*TreeNode
}

func Constructor() Codec {
	return Codec{}
}

// Serializes a tree to a single string.
func (this *Codec) serialize(root *TreeNode) string {
	if root == nil {
		return this.str
	}
	queue := newQueue()
	var slice []string
	queue.add(root)
	for queue.len() > 0 {
		elem := queue.pop()
		if elem != nil {
			slice = append(slice, strconv.Itoa(elem.Val))
			queue.add(elem.Left)
			queue.add(elem.Right)
		} else {
			slice = append(slice, "nil")
		}

	}
	this.str = cutLastNils(strings.Join(slice, ","))
	return this.str
}

type stack struct {
	nodes  []*TreeNode
	length int
}

func newQueue() stack {
	return stack{
		nodes:  []*TreeNode{},
		length: 0,
	}
}
func (q *stack) len() int {
	return q.length

}
func (q *stack) add(elem *TreeNode) {
	q.nodes = append(q.nodes, elem)
	q.length++
}
func (q *stack) pop() *TreeNode {
	if q.length == 0 {
		return nil
	}
	res := q.nodes[0]
	q.nodes = q.nodes[1:len(q.nodes)]
	q.length--
	return res
}

// Deserializes your encoded data to tree.
func (this *Codec) deserialize(data string) *TreeNode {
	if data == "" {
		return nil
	}
	nodes := strings.Split(data, ",")

	if nodes[0] != "nil" {
		val, _ := strconv.Atoi(nodes[0])
		this.root = &TreeNode{Val: val}
		this.lastLayer = []*TreeNode{this.root}
	}
	layer := 1
	for layer < len(nodes) {
		nodes = nodes[layer:]
		layer = layer * 2
		this.buildWithLayer(layer, nodes)

	}
	return this.root
}

func (this *Codec) buildWithLayer(layer int, nodes []string) {
	j := 0
	for i := 0; i < layer && i < len(nodes); i++ {
		val := (nodes)[i]
		par := this.lastLayer[0]

		if i%2 == 1 {
			this.lastLayer = this.lastLayer[1:]
			j++
			par.Right = this.makeNode(val)
		} else {
			par.Left = this.makeNode(val)
		}

	}

}
func (this *Codec) makeNode(in string) *TreeNode {
	if in != "nil" {
		val, _ := strconv.Atoi(in)
		node := &TreeNode{Val: val}
		this.lastLayer = append(this.lastLayer, node)
		return node
	}
	return nil
}

/**
 * Your Codec object will be instantiated and called as such:
 * ser := Constructor();
 * deser := Constructor();
 * data := ser.serialize(root);
 * ans := deser.deserialize(data);
 */
