package main

import "fmt"

func main() {
	obj := Constructor(1)
	param_1 := obj.Next(-1)
	fmt.Println(param_1)
	param_1 = obj.Next(10)
	fmt.Println(param_1)
	param_1 = obj.Next(3)
	fmt.Println(param_1)
	param_1 = obj.Next(5)
	fmt.Println(param_1)
}

type MovingAverage struct {
	arr  []int
	size int
	sum  int
}

func Constructor(size int) MovingAverage {
	return MovingAverage{arr: make([]int, size), size: 0}
}

func (this *MovingAverage) Next(val int) float64 {
	this.size++
	count := (this.size - 1) % len(this.arr)
	old := this.arr[count]
	this.sum = this.sum - old + val
	this.arr[count] = val

	c := len(this.arr)
	if this.size <= len(this.arr) {
		c = this.size
	}
	var res float64
	res = float64(this.sum) / float64(c)
	return res
}

/**
 * Your MovingAverage object will be instantiated and called as such:
 * obj := Constructor(size);
 * param_1 := obj.Next(val);
 */
