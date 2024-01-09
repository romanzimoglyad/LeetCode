package main

import "fmt"

type T struct{}

func main() {
	var t *T
	v := T{}
	t = &v
	test(t)
	test(v)
	fmt.Println(t == nil) // true

	fmt.Println(t.ReturnNum()) // 2

	var n N
	fmt.Println(n == nil)      // true
	fmt.Println(n.ReturnNum()) // panic
	n = t
	fmt.Println(n == nil)      // false
	fmt.Println(n.ReturnNum()) // 2

	test(n)

}

func test(n N) {

}

type N interface {
	ReturnNum() int
	ReturnNum() int
}

func (t *T) ReturnNum() int {
	return 2
}
