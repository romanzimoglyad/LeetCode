package main

import "fmt"

func main() {
	st := CreateDinStack(2)
	st.push(1)
	st.push(2)
	st.push(3)
	st.push(4)
	st.push(5)
	st.push(6)
	fmt.Println(st.popAt(1))
	fmt.Println(st.popAt(1))
	fmt.Println(st.pop())
	fmt.Println(st.pop())
	fmt.Println(st.pop())
	fmt.Println(st)
}

type DinStack struct {
	m  map[int][]int
	i  int
	td int
}

func CreateDinStack(td int) *DinStack {
	return &DinStack{
		m:  make(map[int][]int),
		td: td,
	}
}

func (d *DinStack) push(val int) {
	if len(d.m[d.i]) == d.td {
		d.i++
	}
	d.m[d.i] = append(d.m[d.i], val)
}

func (d *DinStack) pop() int {
	lnth := len(d.m[d.i])
	if lnth <= 1 {
		if lnth == 0 {
			if d.i >= 0 {
				d.i--
				return d.pop()
			}
		}
		resp := d.m[d.i][0]
		d.m[d.i] = []int{}
		if d.i >= 0 {
			d.i--
		}
		return resp
	}

	resp := d.m[d.i][lnth-1]
	d.m[d.i] = d.m[d.i][0 : lnth-1]
	return resp
}

func (d *DinStack) popAt(stackInd int) int {
	if stackInd == d.i {
		return d.pop()
	}
	lnth := len(d.m[stackInd])
	if lnth == 0 {
		return -1
	}
	if lnth == 1 {
		resp := d.m[stackInd][0]
		d.m[stackInd] = []int{}
		return resp
	}

	resp := d.m[stackInd][lnth-1]
	d.m[stackInd] = d.m[stackInd][0 : lnth-1]
	return resp
}
