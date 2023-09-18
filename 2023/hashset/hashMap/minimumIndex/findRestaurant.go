package main

import (
	"fmt"
	"math"
)

func main() {
	arg := []string{"happy", "sad", "good"}
	res := findRestaurant(arg, []string{"sad", "happy", "good"})
	fmt.Println(res)
}

func findRestaurant(list1 []string, list2 []string) []string {
	m1 := make(map[string]int)
	m2 := make(map[string]int)
	m3 := make(map[int][]string)
	min := math.MaxInt
	for l1 := range list1 {
		m1[list1[l1]] = l1
	}
	for l2 := range list2 {
		m2[list2[l2]] = l2
	}
	for ind1, v1 := range m1 {
		if v2, ok := m2[ind1]; ok {
			if v1+v2 <= min {
				min = v1 + v2
			}
			m3[v1+v2] = append(m3[v1+v2], ind1)
		}
	}
	return m3[min]
}
