package main

import "fmt"

func main() {
	l := Constructor()

	t := l.ShouldPrintMessage(1, "foo")
	fmt.Println(t)
	t = l.ShouldPrintMessage(2, "bar")
	fmt.Println(t)
	t = l.ShouldPrintMessage(3, "foo")
	fmt.Println(t)
	t = l.ShouldPrintMessage(8, "bar")
	fmt.Println(t)
	t = l.ShouldPrintMessage(10, "foo")
	fmt.Println(t)
	t = l.ShouldPrintMessage(11, "foo")
	fmt.Println(t)
	t = l.ShouldPrintMessage(12, "foo")
	fmt.Println(t)
	t = l.ShouldPrintMessage(222, "foo")
	fmt.Println(t)
}

type Logger struct {
	m map[string]int
}

func Constructor() Logger {
	m := make(map[string]int)
	return Logger{m: m}
}

func (this *Logger) ShouldPrintMessage(timestamp int, message string) bool {

	if t, ok := this.m[message]; !ok {
		this.m[message] = timestamp + 10
		return true
	} else {
		if timestamp < t {
			return false
		} else {
			this.m[message] = timestamp + 10
			return true
		}
	}
	return false
}

/**
 * Your Logger object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.ShouldPrintMessage(timestamp,message);
 */
