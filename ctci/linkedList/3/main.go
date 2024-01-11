package main

type Node struct {
	val  int
	Next *Node
}

func main() {

}

func remove(cur *Node) {
	cur.val = cur.Next.val
	cur.Next = cur.Next.Next
}
