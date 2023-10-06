package main

import (
	"fmt"
	"math/rand"
)

func main() {
	rs := Constructor()
	fmt.Println(rs.Insert(1))
	fmt.Println(rs.Insert(2))
	fmt.Println(rs.Insert(2))
	fmt.Println(rs.Insert(3))
	fmt.Println(rs.Remove(3))
	fmt.Println(rs.Remove(3))
	fmt.Println(rs.Insert(3))
	fmt.Println(rs.GetRandom())
	fmt.Println(rs.GetRandom())
	fmt.Println(rs.GetRandom())
	fmt.Println(rs.GetRandom())
	fmt.Println(rs.GetRandom())
	fmt.Println(rs.GetRandom())
	fmt.Println(rs.GetRandom())
	fmt.Println(rs.GetRandom())
}

type RandomizedSet struct {
	m   map[int]int
	arr []int
}

func Constructor() RandomizedSet {
	return RandomizedSet{m: make(map[int]int)}
}

func (this *RandomizedSet) Insert(val int) bool {
	if _, ok := this.m[val]; ok {
		return false
	}

	this.arr = append(this.arr, val)
	this.m[val] = len(this.arr) - 1
	return true
}

func (this *RandomizedSet) Remove(val int) bool {
	if _, ok := this.m[val]; ok {
		idx := this.m[val]

		lastItem := this.arr[len(this.arr)-1]
		this.arr[idx] = lastItem
		this.arr = this.arr[0 : len(this.arr)-1]
		this.m[lastItem] = idx
		delete(this.m, val)
		return true
	}
	return false
}

func (this *RandomizedSet) GetRandom() int {

	return this.arr[rand.Intn(len(this.arr))]
}

/**
 * Your RandomizedSet object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Insert(val);
 * param_2 := obj.Remove(val);
 * param_3 := obj.GetRandom();
 */
