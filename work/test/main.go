// Golang program to illustrate the
// concept of the Pointer to struct
package main

import (
	"context"
	"fmt"
	"golang.org/x/time/rate"
	"log"
	"time"
)

func main() {
	limiter := rate.NewLimiter(rate.Limit(1), 10)
	ctx := context.Background()
	//ctx, cancel := context.WithTimeout(ctx, 1*time.Second)
	//defer cancel()
	start := time.Now()
	for i := 0; i < 10; i++ {

		fmt.Println(i, time.Since(start))
		if err := limiter.Wait(ctx); err != nil {
			log.Fatal(err)
		}

	}
	time.Sleep(2 * time.Second)
	for i := 0; i < 10; i++ {

		fmt.Println(i, time.Since(start))
		if err := limiter.Wait(ctx); err != nil {
			log.Fatal(err)
		}

	}
	time.Sleep(2 * time.Second)
	for i := 0; i < 10; i++ {

		fmt.Println(i, time.Since(start))
		if err := limiter.Wait(ctx); err != nil {
			log.Fatal(err)
		}

	}

}

func worker() chan int {
	ch := make(chan int)
	go func() {
		time.Sleep(3 * time.Second)
		ch <- 1
	}()
	return ch
}
