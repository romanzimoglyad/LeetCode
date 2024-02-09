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
func (*MyError) Error() string { return "MyError!" }

func errorHandler(err error) {
	if err != nil {
		fmt.Println("errorHandler function:", err)
	}
}

func main() {
	var err *MyError
	var p error
	err.errorHandler() //вопрос: что выведет?

	errorHandler(err) //вопрос: что выведет?
	errorHandler(p)   //вопрос: что выведет?
}
