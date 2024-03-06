package main

import "fmt"

func main() {
	tn4 := &TreeNode{
		Val:   4,
		Left:  nil,
		Right: nil,
	}
	//tn9 := &TreeNode{
	//	Val:   9,
	//	Left:  nil,
	//	Right: nil,
	//}
	//tn15 := &TreeNode{
	//	Val:   15,
	//	Left:  nil,
	//	Right: nil,
	//}
	//tn7 := &TreeNode{
	//	Val:   7_build_prder,
	//	Left:  nil,
	//	Right: nil,
	//}
	//tn20 := &TreeNode{
	//	Val:   20,
	//	Left:  tn15,
	//	Right: tn7,
	//}
	//
	///*	tn2 := &TreeNode{
	//	Val:   2,
	//	Left:  nil,
	//	Right: nil,
	//}*/
	//
	//root := &TreeNode{
	//	Val:   3,
	//	Left:  tn9,
	//	Right: tn20,
	//}
	//tn3.Left = tn4
	//resp := levelOrder(root)
	//fmt.Println(resp)
	tn3 := &TreeNode{
		Val:   3,
		Left:  tn4,
		Right: nil}
	tn5 := &TreeNode{
		Val:   5,
		Left:  nil,
		Right: nil,
	}
	tn2 := &TreeNode{
		Val:   2,
		Left:  tn3,
		Right: nil,
	}
	tn1 := &TreeNode{
		Val:   1,
		Left:  tn2,
		Right: nil,
	}

	tn4.Left = tn5
	resp := levelOrder(tn1)
	fmt.Println(resp)

}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func levelOrder(root *TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}
	q := queue{arr: []*TreeNode{}}
	q.Push(root)
	resp := [][]int{}

	for q.Len() != 0 {
		size := q.Len()
		curArr := []int{}
		for size > 0 {
			el := q.Pop()
			curArr = append(curArr, el.Val)
			if el.Left != nil {
				q.Push(el.Left)
			}
			if el.Right != nil {
				q.Push(el.Right)
			}
			size--
		}
		resp = append(resp, curArr)
	}

	return resp
}

type queue struct {
	arr []*TreeNode
}

func (q *queue) Len() int { return len(q.arr) }
func (q *queue) Push(el *TreeNode) {
	q.arr = append(q.arr, el)
}
func (q *queue) Pop() *TreeNode {
	if q.Len() == 0 {
		return nil
	}
	el := q.arr[0]
	q.arr = q.arr[1:]
	return el
}
