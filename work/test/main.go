// Golang program to illustrate the
// concept of the Pointer to struct
package main

import (
	"fmt"
	"time"
)

func main() {
	travelPointsClientStartTime := time.Now()
	time.Sleep(1 * time.Second)
	fmt.Println(time.Since(travelPointsClientStartTime))
}

func worker() chan int {
	ch := make(chan int)
	go func() {
		time.Sleep(3 * time.Second)
		ch <- 1
	}()
	return ch
}
