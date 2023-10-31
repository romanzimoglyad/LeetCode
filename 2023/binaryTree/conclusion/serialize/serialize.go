package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	tn5 := &TreeNode{
		Val:   5,
		Left:  nil,
		Right: nil,
	}
	tn6 := &TreeNode{
		Val:   6,
		Left:  nil,
		Right: nil,
	}
	tn7 := &TreeNode{
		Val:   7,
		Left:  nil,
		Right: nil,
	}
	tn4 := &TreeNode{
		Val:   4,
		Left:  nil,
		Right: nil,
	}
	tn3 := &TreeNode{
		Val:   3,
		Left:  tn6,
		Right: tn7}

	tn2 := &TreeNode{
		Val:   2,
		Left:  tn4,
		Right: tn5,
	}
	tn1 := &TreeNode{
		Val:   1,
		Left:  tn2,
		Right: tn3,
	}
	fmt.Println(tn1)
	ser := Constructor()
	//deser := Constructor()
	data := ser.serialize(tn1)
	//tn33 := &TreeNode{
	//	Val:   3,
	//	Left:  nil,
	//	Right: nil}
	//
	//tn22 := &TreeNode{
	//	Val:   2,
	//	Left:  nil,
	//	Right: nil,
	//}
	//tn11 := &TreeNode{
	//	Val:   1,
	//	Left:  tn22,
	//	Right: tn33,
	//}
	//
	//data := ser.serialize(tn11)
	r := ser.deserialize(data)
	fmt.Println(r)
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type Codec struct {
	str   string
	elems []string
}

func Constructor() Codec {
	return Codec{}
}

// Serializes a tree to a single string.
func (this *Codec) serialize(root *TreeNode) string {
	if root == nil {
		return ""
	}
	return this.tr1(root)
}

// Deserializes your encoded data to tree.
func (this *Codec) deserialize(data string) *TreeNode {
	if data == "" {
		return nil
	}
	elems := strings.Split(data, ",")
	this.elems = elems
	root := this.deserializeHelper()
	return root
}

func (this *Codec) deserializeHelper() *TreeNode {
	if len(this.elems) == 0 {
		return nil
	}
	elem := this.elems[0]
	this.elems = this.elems[1:]
	if elem == "nil" {
		return nil
	}
	node := makeNode(elem)

	node.Left = this.deserializeHelper()
	node.Right = this.deserializeHelper()
	return node
}

func makeNode(val string) *TreeNode {
	v, _ := strconv.Atoi(val)
	return &TreeNode{
		Val:   v,
		Left:  nil,
		Right: nil,
	}

}

func (this *Codec) tr1(root *TreeNode) string {
	if root == nil {
		this.str += ",nil"
	} else {
		if this.str == "" {
			this.str += strconv.Itoa(root.Val)
		} else {
			this.str += "," + strconv.Itoa(root.Val)
		}
		this.tr1(root.Left)
		this.tr1(root.Right)
	}
	return this.str
}

/**
 * Your Codec object will be instantiated and called as such:
 * ser := Constructor();
 * deser := Constructor();
 * data := ser.serialize(root);
 * ans := deser.deserialize(data);
 */
