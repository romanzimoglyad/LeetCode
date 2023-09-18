package main

import (
	"fmt"
)

func main() {

	obj := Constructor()
	obj.Put(1, 1)
	obj.Put(2, 2)
	obj.Put(2, 1)
	t := obj.Get(1)
	obj.Remove(2)
	t = obj.Get(2)

	fmt.Println(t)
}

var bucketsNum int = 100000

type Pair struct {
	key   int
	value int
}
type MyHashMap struct {
	m [][]Pair
}

func Constructor() MyHashMap {
	return MyHashMap{
		m: make([][]Pair, bucketsNum),
	}
}

func (this *MyHashMap) Put(key int, value int) {

	ind := getNumBucket(key)
	if this.Get(key) != -1 {
		for i := 0; i < len(this.m[ind]); i++ {
			if this.m[ind][i].key == key {
				this.m[ind][i].value = value
				return
			}
		}
	}
	if this.m[ind] == nil {
		this.m[ind] = make([]Pair, 0)
	}
	this.m[ind] = append(this.m[ind], Pair{
		key:   key,
		value: value,
	})
}

func (this *MyHashMap) Get(key int) int {
	ind := getNumBucket(key)
	if this.m[ind] != nil {
		for i := 0; i < len(this.m[ind]); i++ {
			if this.m[ind][i].key == key {
				return this.m[ind][i].value
			}
		}
	}
	return -1
}

func (this *MyHashMap) Remove(key int) {
	ind := getNumBucket(key)
	if this.m[ind] != nil {
		for i := 0; i < len(this.m[ind]); i++ {
			if this.m[ind][i].key == key {
				this.m[ind] = append(this.m[ind][0:i], this.m[ind][i+1:]...)
			}
		}
	}
}

func getNumBucket(val int) int {
	return val % bucketsNum
}
