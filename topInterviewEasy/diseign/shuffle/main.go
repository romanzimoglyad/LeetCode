package main

import (
	"fmt"
	"math/rand"
)

func main() {
	s := Constructor([]int{1, 2, 3, 4, 5})
	fmt.Println(s.Shuffle())
	fmt.Println(s.Shuffle())
	fmt.Println(s.Shuffle())
	fmt.Println(s.Shuffle())
	fmt.Println(s.Reset())
}

type Solution struct {
	baseArr []int
	newArr  []int
	Seed    int
}

func Constructor(nums []int) Solution {
	new := make([]int, len(nums))
	copy(new, nums)
	return Solution{
		baseArr: nums,
		newArr:  new,
		Seed:    len(nums),
	}
}

func (this *Solution) Reset() []int {
	return this.baseArr
}

func (this *Solution) Shuffle() []int {
	for i := range this.baseArr {
		r := i + rand.Intn(len(this.baseArr)-i)
		this.newArr[i], this.newArr[r] = this.newArr[r], this.newArr[i]
	}
	return this.newArr
}

/**
 * Your Solution object will be instantiated and called as such:
 * obj := Constructor(nums);
 * param_1 := obj.Reset();
 * param_2 := obj.Shuffle();
 */
