package main

import "fmt"

//["KthLargest","add","add","add","add","add","add","add","add","add","add","add","add","add","add","add","add","add","add","add","add","add","add","add","add","add","add","add"]
//[[7_build_prder,[-10,1,3,1,4,10,3,9,4,5,1]],[3],[2],[3],[1],[2],[4],[5],[5],[6],[7_build_prder],[7_build_prder],[8],[2],[3],[1],[1],[1],[10],[11],[5],[6],[2],[4],[7_build_prder],[8],[5],[6]]

func main() {
	tree := Constructor(3, []int{4, 5, 8, 2})
	fmt.Println(tree.Add(3))
	fmt.Println(tree.Add(5))
	fmt.Println(tree.Add(10))
	fmt.Println(tree.Add(9))
	fmt.Println(tree.Add(4))

}

type KthLargest struct {
	k    int
	root *node
}

type node struct {
	val   int
	left  *node
	right *node
	k     int
}

func Constructor(k int, nums []int) KthLargest {
	var root *node

	for _, v := range nums {
		root = addNode(root, v)
	}

	return KthLargest{k: k,
		root: root}
}

func addNode(root *node, val int) *node {
	if root == nil {
		root = &node{val: val, k: 1}
		return root
	}
	if root.val >= val {
		root.left = addNode(root.left, val)
	}
	if root.val < val {
		root.right = addNode(root.right, val)

	}
	root.k++

	return root
}

func (this *KthLargest) Add(val int) int {
	this.root = addNode(this.root, val)
	n, _ := findK(this.root, val, this.k)
	return n.val
}

func findK(root *node, val int, k int) (*node, int) {
	if root == nil {
		return nil, k
	}
	if root.k == k {
		if root.left == nil {
			return root, 0
		}
		return findK(root.left, val, k-1)
	}
	if root.k < k {
		return nil, k - root.k
	}

	if root.k > k {
		n, newK := findK(root.right, val, k)
		if n != nil {
			return n, 0
		}
		if newK == 1 {
			return root, 0
		}
		return findK(root.left, val, newK-1)
	}
	return nil, 0
}
