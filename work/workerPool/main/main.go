package main

import (
	"context"
	"fmt"
	"github.com/romanzimoglyad/LeetCode/work/workerPool"
	"time"
)

func main() {

	wp := workerPool.NewWorkerPool(5)
	_, cancel := context.WithTimeout(context.Background(), time.Second*1115)
	defer cancel()

	tasks := []workerPool.Task{}

	for i := 0; i < 1000; i++ {
		tasks = append(tasks, workerPool.Task{I: i})
	}
	resp := wp.AddTasks(tasks)

	/*	for i := 0; i < 20; i++ {

			go func() {
				time.Sleep(time.Second * 2)
				wp.AddTasks(tasks)
			}()
		}
	*/
	go func() {
		for val := range resp {
			fmt.Println("result: %v", val.Resp)
		}
	}()

	time.Sleep(time.Second * 50)

}
