package main

import "fmt"

func main() {
	fmt.Println(dailyTemperatures([]int{7, 5, 7, 6, 7, 84, 2, 1, 66, 666}))
}

func dailyTemperatures(temperatures []int) []int {
	st := []int{}

	res := make([]int, len(temperatures))

	for key, value := range temperatures {

		for len(st) > 0 && temperatures[st[len(st)-1]] < value {
			v := st[len(st)-1]
			st = st[0 : len(st)-1]

			res[v] = key - v
		}
		st = append(st, key)

	}

	return res
}
