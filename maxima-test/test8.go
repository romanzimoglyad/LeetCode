package main

import (
	"fmt"
	"sync"
)

var num int

func main() {
	// Какой будет результат выполнения приложения
	w := sync.WaitGroup{}
	for i := 0; i < 1000; i++ {
		w.Add(1)
		go func() {
			defer w.Done()
			num = i
		}()
	}
	w.Wait()
	fmt.Printf("NUM is %d", num)
}
