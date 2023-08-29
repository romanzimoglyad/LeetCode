// Golang program to illustrate the
// concept of the Pointer to struct
package main

import (
	"fmt"
	"time"
)

func main() {
	var p *int
	t := 1
	p = &t
	k := &p
	*p = 4
	fmt.Println(t)
	fmt.Println(*p)
	fmt.Println(k)
	b := 3
	v := &b
	*k = v
	fmt.Println(k)
	fmt.Println(*p)
}

func worker() chan int {
	ch := make(chan int)
	go func() {
		time.Sleep(3 * time.Second)
		ch <- 1
	}()
	return ch
}
