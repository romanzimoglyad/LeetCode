package main

import "fmt"

func main() {

	obj := Constructor()
	obj.Add(1)
	obj.Add(1111)
	obj.Add(111111)
	obj.Remove(2)
	r := obj.Contains(111111)
	fmt.Println(r)
}

type MyHashSet struct {
	m []map[int]struct{}
}

func Constructor() MyHashSet {
	m := make([]map[int]struct{}, 1000)
	return MyHashSet{m: m}
}

func (this *MyHashSet) Add(key int) {
	bucketNumber := hashSetFunc(key)
	if this.m[bucketNumber] == nil {
		this.m[bucketNumber] = make(map[int]struct{})
	}
	this.m[bucketNumber][key] = struct{}{}
}

func (this *MyHashSet) Remove(key int) {
	bucketNumber := hashSetFunc(key)

	delete(this.m[bucketNumber], key)
}

func (this *MyHashSet) Contains(key int) bool {
	bucketNumber := hashSetFunc(key)
	if _, ok := this.m[bucketNumber][key]; ok {
		return ok
	}
	return false
}

func hashSetFunc(val int) int {
	return val % 1000
}
