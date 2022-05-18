package main

import (
	"fmt"
	"sync"
)

func main() {
	t := struct {
		s string
		i int
		b bool
	}{}
	w := new(sync.WaitGroup)
	w.Add(1)
	go func() {
		defer w.Done()
		t.s = "sasd"
	}()
	w.Add(1)
	go func() {
		defer w.Done()
		t.s = "sa123123sd"
	}()
	w.Add(1)
	go func() {
		defer w.Done()
		t.s = "12312312"
	}()
	w.Add(1)
	go func() {
		defer w.Done()
		t.i = 2
	}()
	w.Add(1)
	go func() {
		defer w.Done()
		t.b = true
	}()
	w.Wait()
	fmt.Printf("%v", t)
}
