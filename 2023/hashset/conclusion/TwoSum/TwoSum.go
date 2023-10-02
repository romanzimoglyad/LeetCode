package main

import "fmt"

func main() {
	twoSum := Constructor()
	twoSum.Add(1)
	twoSum.Add(3)
	twoSum.Add(5)
	fmt.Println(twoSum.Find(4))
	fmt.Println(twoSum.Find(7))
}

type TwoSum struct {
	array []int
}

func Constructor() TwoSum {
	return TwoSum{array: make([]int, 0)}
}

func (this *TwoSum) Add(number int) {
	this.array = append(this.array, number)
}

func (this *TwoSum) Find(value int) bool {
	m := make(map[int]struct{})
	for i := 0; i < len(this.array); i++ {

		if _, ok := m[this.array[i]]; ok {
			return true
		} else {
			m[value-this.array[i]] = struct{}{}
		}
	}

	return false
}

/**
 * Your TwoSum object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Add(number);
 * param_2 := obj.Find(value);
 */
