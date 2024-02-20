package main

import "fmt"

func main() {
	//r10 := &TreeNode{Val: 10}
	//
	//r2 := &TreeNode{Val: 2}
	//r9 := &TreeNode{Val: 9}
	//r5 := &TreeNode{Val: 5}
	//r6 := &TreeNode{Val: 6}
	//r7 := &TreeNode{Val: 7}
	//r8 := &TreeNode{Val: 8}
	//r10.Left = r5
	//r5.Right = r9
	//r9.Left = r7
	//r5.Left = r2
	//r7.Right = r8
	//r7.Left = r6
	//c := Constructor(r10)
	//
	//fmt.Println(c.Next())
	//fmt.Println(c.Next())
	//fmt.Println(c.Next())
	//fmt.Println(c.Next())
	//fmt.Println(c.Next())
	//fmt.Println(c.Next())
	r3 := &TreeNode{Val: 3}

	r15 := &TreeNode{Val: 15}
	r9 := &TreeNode{Val: 9}
	r20 := &TreeNode{Val: 20}

	r7 := &TreeNode{Val: 7}

	r7.Left = r3
	r15.Right = r20
	r7.Right = r15
	r15.Left = r9
	c := Constructor(r7)

	fmt.Println(c.Next())
	fmt.Println(c.Next())
	fmt.Println(c.HasNext())

	fmt.Println(c.Next())
	fmt.Println(c.HasNext())
	fmt.Println(c.Next())
	fmt.Println(c.HasNext())
	fmt.Println(c.Next())
	fmt.Println(c.HasNext())
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type BSTIterator struct {
	root *TreeNode

	prev *TreeNode
}

func Constructor(root *TreeNode) BSTIterator {
	leftMost := root
	for leftMost.Left != nil {
		leftMost = leftMost.Left
	}
	new := &TreeNode{
		Val: -1,
	}
	new.Left = leftMost
	return BSTIterator{root: root,
		prev: new}
}

func (this *BSTIterator) Next() int {
	res := this.inorderSuccessor()
	this.prev = res
	return res.Val
}

func (this *BSTIterator) inorderSuccessor() *TreeNode {

	cur := this.root
	var parent *TreeNode
	for cur != nil {
		if this.prev.Val >= cur.Val {
			cur = cur.Right
		} else {
			parent = cur
			cur = cur.Left
		}
	}

	return parent
}

func (this *BSTIterator) HasNext() bool {
	return this.inorderSuccessor() != nil
}

/**
 * Your BSTIterator object will be instantiated and called as such:
 * obj := Constructor(root);
 * param_1 := obj.Next();
 * param_2 := obj.HasNext();
 */
