package main

import "fmt"

func main() {
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	resp := createTree(arr)
	fmt.Println(resp)
}

type Node struct {
	val   int
	left  *Node
	right *Node
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

// 1 2 3 4 5 6 7_build_prder 8 9 10
