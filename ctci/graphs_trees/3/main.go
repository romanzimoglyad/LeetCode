package main

import "fmt"

func main() {
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	resp := createTree(arr)
	rr := getNodes(resp)
	fmt.Println(rr)
}
func createTree(arr []int) *Node {
	l := len(arr)
	if l == 0 {
		return nil
	}
	if l == 1 {
		return &Node{val: arr[0]}
	}
	mid := l / 2
	root := &Node{
		val:   arr[mid],
		left:  createTree(arr[0:mid]),
		right: createTree(arr[mid+1:]),
	}
	return root
}

func getNodes(root *Node) []*listNode {
	answer := []*listNode{}
	q := Queue{arr: make([]*Node, 0, 0)}
	if root == nil {
		return []*listNode{}
	}
	q.add(root)
	for q.len() != 0 {
		l := q.len()
		dummy := &listNode{}
		cur := dummy
		for i := 0; i < l; i++ {
			newEl := q.get()
			cur.next = &listNode{
				val: newEl,
			}
			cur = cur.next
			if newEl.left != nil {
				q.add(newEl.left)
			}
			if newEl.right != nil {
				q.add(newEl.right)
			}
		}
		answer = append(answer, dummy.next)
	}
	return answer
}

type listNode struct {
	val  *Node
	next *listNode
}

type Node struct {
	val   int
	left  *Node
	right *Node
}

type Queue struct {
	arr []*Node
}

func (q *Queue) add(el *Node) {
	q.arr = append(q.arr, el)
}

func (q *Queue) get() *Node {
	l := len(q.arr)
	if l == 0 {
		return nil
	}
	el := q.arr[0]
	q.arr = q.arr[1:]
	return el
}
func (q *Queue) len() int {
	return len(q.arr)
}
