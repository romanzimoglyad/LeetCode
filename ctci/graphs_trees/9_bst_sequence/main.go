package main

import "fmt"

func main() {
	n1 := &Node{val: 1}
	n2 := &Node{val: 2}
	n3 := &Node{val: 3}

	n2.left = n1
	n2.right = n3

	r := buildArr(n2)
	fmt.Println(r)

}

var arr []int

type Node struct {
	val   int
	left  *Node
	right *Node
}

func buildArr(root *Node) [][]int {
	if root == nil {
		return [][]int{}
	}
	left := buildArr(root.left)
	right := buildArr(root.right)
	var resp [][]int

	if len(left) == 0 && len(right) == 0 {
		return [][]int{{root.val}}
	}

	for i := 0; i < len(left); i++ {
		for j := 0; j < len(right); j++ {
			var respElement1 []int
			respElement1 = append(respElement1, root.val)
			respElement1 = append(respElement1, left[i]...)
			respElement1 = append(respElement1, right[i]...)
			resp = append(resp, respElement1)
			var respElement2 []int
			respElement2 = append(respElement2, root.val)
			respElement2 = append(respElement2, right[i]...)
			respElement2 = append(respElement2, left[i]...)
			resp = append(resp, respElement2)
		}
	}
	return resp
}

//
//func buildArr2(root *Node) []int {
//	elems := [][]int{}
//	queue := []*Node{}
//	if root == nil {
//		return []int{}
//	}
//	queue = append(queue, root)
//	j := 0
//	for l := len(queue); l != 0; {
//		for i := 0; i < l; i++ {
//			el := queue[0]
//			queue = queue[1:]
//			elems[j] = append(elems[j], el.val)
//			if el.left != nil {
//				queue = append(queue, el.left)
//			}
//			if el.right != nil {
//				queue = append(queue, el.right)
//			}
//		}
//		j++
//	}
//
//}
//
//func buildAllArrays(in []int) [][]int {
//
//}
//
//func inOrder(root *Node) {
//	if root == nil {
//		return
//	}
//	inOrder(root.left)
//	arr = append(arr, root.val)
//	inOrder(root.right)
//}
