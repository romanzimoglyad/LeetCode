package main

import (
	"fmt"
)

type SyncInt struct {
	counter int
	c       chan struct{}
}

func main() {
	var counter SyncInt
	counter.c = make(chan struct{})
	for i := 0; i < 1000; i++ {
		go func() {
			counter.c <- struct{}{}
			counter.counter++
		}()
		<-counter.c
	}

	fmt.Println(counter.counter)
}
