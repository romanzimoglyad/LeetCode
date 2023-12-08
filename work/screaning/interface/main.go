package main

import "fmt"

/**
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
	var p error
	err.errorHandler() //вопрос: что выведет?

	errorHandler(err) //вопрос: что выведет?
	errorHandler(p)   //вопрос: что выведет?
}

*/

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
	err.errorHandler() //вопрос: что выведет?
	// Ничего, так как err = nil (переменная не инициализирована)
	errorHandler(err) //вопрос: что выведет?
	// "errorHandler function: <nil>"
	// Интерфейс - кортеж из значения и конкретного типа
	// В данном случае тип - *MyError, значение - nil, но сам интерфейс - не nil
	err = &MyError{}
	err.errorHandler() //вопрос: что выведет?
	// Error Method: MyError!
	errorHandler(err) //вопрос: что выведет?
	// errorHandler function: MyError!
}
