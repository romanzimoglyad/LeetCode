package main

import (
	"fmt"
)

type MyError struct{}

func (err *MyError) errorHandler() {
	if err != nil {
		fmt.Println("Error Method:", err)
	}
}
func (MyError) Error() string { return "MyError!" }

func errorHandler(err error) {
	if err != nil {
		fmt.Println("errorHandler function:", err)
	}
}

func main() {
	var err *MyError
	err.errorHandler()
	errorHandler(err)

	err = &MyError{}
	err.errorHandler()
	errorHandler(err)
}
