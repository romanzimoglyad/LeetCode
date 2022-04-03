package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {

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
		return ""
	}
	this.ser(root)
	return this.str
}

// Serializes a tree to a single string.
func (this *Codec) ser(root *TreeNode) {
	if root == nil {
		this.str += ",nil"
	} else {
		if this.str == "" {
			this.str += strconv.Itoa(root.Val)
		} else {
			this.str += "," + strconv.Itoa(root.Val)
		}

		this.ser(root.Left)
		this.ser(root.Right)
	}
}

// Deserializes your encoded data to tree.
func (this *Codec) deserialize(data string) *TreeNode {
	if data == "" {
		return nil
	}
	nodes := strings.Split(data, ",")
	return this.des(&nodes)
}
func (this *Codec) des(nodes *[]string) *TreeNode {
	val := (*nodes)[0]
	*nodes = (*nodes)[1:]
	if val == "nil" {
		return nil
	}

	root := this.makeNode(val)
	root.Left = this.des(nodes)
	root.Right = this.des(nodes)
	return root

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
