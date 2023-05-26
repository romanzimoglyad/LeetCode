// Golang program to illustrate the
// concept of the Pointer to struct
package main

import (
	"fmt"
	"time"
)

func main() {
	m := make(map[string]int)
	m["key"] = 1
	v := m["key"]

	str := "asd"
	p := &m["key"]
	t := &str

	fmt.Println(p)
	fmt.Println(t)
}

func worker() chan int {
	ch := make(chan int)
	go func() {
		time.Sleep(3 * time.Second)
		ch <- 1
	}()
	return ch
}
