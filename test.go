package main

import (
	"context"
	"fmt"
	"time"
)

func firstMicroservice(ctx context.Context) {
	// Simulating work that takes some time

	select {
	case <-secondMicroservice(ctx):
		fmt.Println(time.Now(), "Second microservice completed successfully")
	case <-ctx.Done():
		fmt.Println(time.Now(), "First microservice canceled due to context deadline exceeded")
	}
}

func secondMicroservice(ctx context.Context) <-chan struct{} {
	// Simulating work that takes some time
	resp := make(chan struct{})
	newCTXWithDeadline, _ := context.WithTimeout(ctx, 1*time.Second) //nolint:gomnd // ok

	go func() {
		select {
		case <-time.After(6 * time.Second):
			//fmt.Println("asdasd")
			resp <- struct{}{}
		case <-newCTXWithDeadline.Done():
			fmt.Println(time.Now(), "Second microservice canceled due to context deadline exceeded")
		}
	}()

	return resp
}

func main() {
	fmt.Println(time.Now(), "start")
	ctx, cancel := context.WithDeadline(context.Background(), time.Now().Add(5*time.Second))
	defer cancel()

	// Call the first microservice with the original context
	firstMicroservice(ctx)

	//// Change the deadline of the context
	//newCtx, newCancel := context.WithDeadline(ctx, time.Now().Add(2*time.Second))
	//defer newCancel()
	//
	//// Call the second microservice with the new context
	//secondMicroservice(newCtx)
	time.Sleep(10 * time.Second)
}
