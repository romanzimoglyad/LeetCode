package main

import (
	"fmt"
	"sync"
)

func main() {
	// Какой будет результат выполнения приложения
	ch := make(chan int)
	wg := &sync.WaitGroup{}
	wg.Add(3)
	for i := 0; i < 3; i++ {
		go func(idx int) {

			ch <- (idx + 1) * 2
			fmt.Println("here")
			wg.Done()
		}(i)
	}
	fmt.Printf("result: %d\n", <-ch)
	fmt.Printf("result: %d\n", <-ch)
	fmt.Printf("result: %d\n", <-ch)
	wg.Wait()
}
