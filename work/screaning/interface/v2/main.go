package main

import "fmt"

type Foo struct{}

func (f *Foo) A() {}
func (f *Foo) B() {}
func (f *Foo) C() {}

type AB interface {
	A()
	B()
}

type BC interface {
	B()
	C()
}

func main() {

	t := "<font color=textNegative>⁺¹</font>"
	fmt.Println(t)
	tt := `\U+2081`
	fmt.Println(tt)

}
