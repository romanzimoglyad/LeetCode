package main

import (
	"fmt"
	"github.com/emirpasic/gods/lists/singlylinkedlist"
)

func main() {
	s := Shelter[Animal]{l: singlylinkedlist.New()}
	s.Enque(Cat{})
	s.Enque(Cat{})
	s.Enque(Dog{})
	el := s.GetDog()
	fmt.Println(el)
	el = s.Get()
	fmt.Println(el)
	el = s.Get()
	fmt.Println(el)
	el = s.Get()
	fmt.Println(el)
	el = s.GetDog()
	fmt.Println(el)
}

type Animal interface {
	isDog() bool
	isCat() bool
}

type Cat struct {
}

func (c Cat) isDog() bool { return false }
func (c Cat) isCat() bool { return true }

type Dog struct {
}

func (c Dog) isDog() bool { return true }
func (c Dog) isCat() bool { return false }

type Shelter[K Animal] struct {
	l *singlylinkedlist.List
}

func (s Shelter[K]) Enque(el K) {
	s.l.Add(el)
}

func (s Shelter[K]) Get() K {
	el, ok := s.l.Get(0)
	if !ok {
		panic("asd")
	}
	p := el.(K)
	s.l.Remove(0)
	return p
}

func (s Shelter[K]) GetDog() Animal {
	return s.getByType("dog")
}
func (s Shelter[K]) GetCat() Animal {
	return s.getByType("cat")
}

func (s Shelter[K]) getByType(t string) Animal {
	index := 0
	for index < s.l.Size() {
		el, ok := s.l.Get(index)
		if !ok {
			return nil
		}
		var an Animal
		if t == "cat" {
			an, ok = el.(Cat)
			if !ok {
				index++
				continue
			}
		}
		if t == "dog" {
			an, ok = el.(Dog)
			if !ok {
				index++
				continue
			}
		}

		s.l.Remove(index)
		return an
	}
	return nil
}
