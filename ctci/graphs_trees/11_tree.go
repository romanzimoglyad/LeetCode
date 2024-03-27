package main

import (
	"fmt"
	"math/rand"
)

func main() {
	n5 := &TreeNode{Val: 5}
	n3 := &TreeNode{Val: 3}
	n6 := &TreeNode{Val: 6}
	n2 := &TreeNode{Val: 2}
	n4 := &TreeNode{Val: 4}
	n7 := &TreeNode{Val: 7}

	n5.Left = n3
	n5.Right = n6
	n3.Left = n2
	n3.Right = n4
	n6.Right = n7
	t := Tree{root: n5}
	a := t.Search(71)
	fmt.Println(a)
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
	size  int
}

type Tree struct {
	root *TreeNode
}

func (t *Tree) Insert(val int) {

	insert(t.root, val)
}

func insert(t *TreeNode, val int) {
	if t == nil {
		t = &TreeNode{Val: val}
		return
	}
	cur := t
	for cur != nil {
		if cur.Val > val {
			if cur.Left == nil {
				cur.Left = &TreeNode{Val: val, size: 1}

			} else {
				cur = cur.Left
			}
		} else {

			if cur.Right == nil {
				cur.Right = &TreeNode{Val: val}

			}
			cur = cur.Right
		}
	}
	cur.size++
}

func (t *TreeNode) Delete(val int) {

	delete(t, val)
}

func delete(t *TreeNode, val int) *TreeNode {
	if t == nil {
		return nil
	}
	if t.Val == val {
		if t.Left == nil && t.Right == nil {
			return nil
		}
		if t.Right != nil {
			suc := successor(t)
			t.Val = suc.Val
			t.Right = delete(t.Right, suc.Val)
		} else {
			pred := predeccessor(t)
			t.Val = pred.Val
			t.Left = delete(t.Left, pred.Val)
		}
	}
	if t.Val > val {
		t.Left = delete(t.Left, val)

	} else if t.Val > val {
		t.Right = delete(t.Right, val)
	}
	t.size--

	return t
}

func successor(t *TreeNode) *TreeNode {
	cur := t.Right
	for cur.Left != nil {
		cur = cur.Left
	}
	return cur
}
func predeccessor(t *TreeNode) *TreeNode {
	cur := t.Left
	for cur.Right != nil {
		cur = cur.Right
	}
	return cur
}

func (t *Tree) Search(val int) *TreeNode {
	return search(t.root, val)
}

func search(t *TreeNode, val int) *TreeNode {
	if t == nil {
		return nil
	}
	if t.Val == val {
		return t
	}
	if t.Val > val {
		return search(t.Left, val)
	}

	return search(t.Right, val)
}

func (t *TreeNode) GetRandom() *TreeNode {
	r := rand.Intn(t.size)
	var leftSize int
	if t.Left != nil {
		leftSize = t.Left.size
	}

	if r < leftSize {
		return t.Left.GetRandom()
	} else if r == leftSize {
		return t
	} else {
		return t.Right.GetRandom()
	}
	return nil
}
