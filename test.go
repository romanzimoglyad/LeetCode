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
	var time1 time.Time
	var time2 time.Time
	time1, _ = time.Parse(time.RFC3339, "2023-07-21T01:00:05Z")
	time2, _ = time.Parse(time.RFC3339, "2023-07-22T23:59:05Z")

	timeWithoutTime := time.Date(time1.Year(), time1.Month(), time1.Day(), 0, 0, 0, 0, time1.Location())
	dateWithoutTime2 := time.Date(time2.Year(), time2.Month(), time2.Day(), 0, 0, 0, 0, time2.Location())
	days := dateWithoutTime2.Sub(timeWithoutTime).Hours() / 24
	fmt.Println(days)

	time1, _ = time.Parse(time.RFC3339, "2023-07-21T01:00:05Z")
	time2, _ = time.Parse(time.RFC3339, "2023-07-22T01:59:05Z")
	timeWithoutTime = time.Date(time1.Year(), time1.Month(), time1.Day(), 0, 0, 0, 0, time1.Location())
	dateWithoutTime2 = time.Date(time2.Year(), time2.Month(), time2.Day(), 0, 0, 0, 0, time2.Location())
	days = dateWithoutTime2.Sub(timeWithoutTime).Hours() / 24
	fmt.Println(days)

	time1, _ = time.Parse(time.RFC3339, "2023-07-21T23:00:05Z")
	time2, _ = time.Parse(time.RFC3339, "2023-07-23T01:59:05Z")
	timeWithoutTime = time.Date(time1.Year(), time1.Month(), time1.Day(), 0, 0, 0, 0, time1.Location())
	dateWithoutTime2 = time.Date(time2.Year(), time2.Month(), time2.Day(), 0, 0, 0, 0, time2.Location())
	days = dateWithoutTime2.Sub(timeWithoutTime).Hours() / 24
	fmt.Println(days)

}
