package main

import "fmt"

func main() {

	res := levelOrder(nil)
	fmt.Println(res)
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
	res := make([][]int, 0, 0)
	q := queue{}
	q.add(root)
	q.add(nil)
	i := 0
	tmp := []int{}
	for !q.empty() {

		el := q.top()
		if el == nil {
			q.pop()
			res = append(res, tmp)
			tmp = []int{}
			i++
			if !q.empty() {
				q.add(nil)
			}
			continue
		}
		tmp = append(tmp, el.Val)
		if el.Left != nil {
			q.add(el.Left)
		}
		if el.Right != nil {
			q.add(el.Right)
		}
		q.pop()

	}
	return res
}

type queue struct {
	val []*TreeNode
}

func (q *queue) add(v *TreeNode) {
	q.val = append(q.val, v)
}
func (q *queue) top() *TreeNode {
	if !q.empty() {
		return q.val[0]
	}
	return nil
}
func (q *queue) pop() {
	q.val = q.val[1:]
}
func (q *queue) empty() bool {
	return len(q.val) == 0
}
