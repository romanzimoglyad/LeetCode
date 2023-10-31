package main

import "fmt"

func main() {
	r := Constructor(1)
	fmt.Println(r.Next(41))
	fmt.Println(r.Next(0))

}

type MovingAverage struct {
	arr  []int
	size int ``
}

func Constructor(size int) MovingAverage {
	return MovingAverage{
		arr:  []int{},
		size: size,
	}
}

func (this *MovingAverage) Next(val int) float64 {
	this.arr = append(this.arr, val)
	if len(this.arr) <= this.size {
		res := 0
		for i := range this.arr {
			res += this.arr[i]
		}
		return float64(res) / float64(len(this.arr))
	}
	res := 0
	for i := len(this.arr) - this.size; i <= len(this.arr)-1; i++ {
		res += this.arr[i]

	}

	return float64(res) / float64(this.size)
}

/**
 * Your MovingAverage object will be instantiated and called as such:
 * obj := Constructor(size);
 * param_1 := obj.Next(val);
 */
