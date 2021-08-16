package main

import (
	"fmt"
	"os"
	"os/signal"
)

func main() {
	testCh := make(chan string)

	go func(chan string) {
		t := <-testCh
		fmt.Println("1_findMaxConsecutiveOnes" + t)
	}(testCh)
	go func(chan string) {
		t := <-testCh
		fmt.Println("2_findEvenNumbers" + t)
	}(testCh)
	go func(chan string) {
		t := <-testCh
		fmt.Println("3_SortedSquares" + t)
	}(testCh)
	close(testCh)
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, os.Interrupt)
	_ = <-sigs
	fmt.Println("Получен сигнал:. Остановка приложения.")
}
