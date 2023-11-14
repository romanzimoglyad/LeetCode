package main

import (
	"container/list"
	"fmt"
	"strconv"
	"strings"
)

func main() {
	res := openLock([]string{"0201", "0101", "0102", "1212", "2002"}, "0202")
	fmt.Println(res)
	res = openLock([]string{"8888"}, "0009")
	fmt.Println(res)
	res = openLock([]string{"8887", "8889", "8878", "8898", "8788", "8988", "7888", "9888"}, "8888")
	fmt.Println(res)
	res = openLock([]string{"0000"}, "8888")
	fmt.Println(res)
	fmt.Println(toArr("123"))
}
func toArr(s string) [4]byte {
	return [4]byte{s[0], s[1], s[2], s[3]}
}

var moves = [][4]int{
	{1, 0, 0, 0},
	{-1, 0, 0, 0},
	{0, 1, 0, 0},
	{0, -1, 0, 0},
	{0, 0, 1, 0},
	{0, 0, -1, 0},
	{0, 0, 0, 1},
	{0, 0, 0, -1}}

var final = [4]int{0, 0, 0, 0}

func openLock(deadends []string, target string) int {

	ddns := splitDeadends(deadends)
	fmt.Println(ddns)
	tgt := split(target)
	if tgt == final && !inDeadends(ddns, tgt) {
		return 0
	}
	q := queue{}
	q.Push(&tgt)
	count := 1
	m := make(map[[4]int]struct{})
	for q.Len() > 0 {
		size := q.Len()
		for i := 0; i < size; i++ {
			combination := q.Pop()
			for _, move := range moves {
				newCombination := [4]int{}
				for i := range move {
					newCombination[i] = combination[i] + move[i]
					if newCombination[i] > 9 {
						newCombination[i] = 0
					} else if newCombination[i] < 0 {
						newCombination[i] = 9
					}
				}
				if newCombination == final && !inDeadends(ddns, newCombination) {
					return count
				}
				if _, ok := m[newCombination]; !ok && !inDeadends(ddns, newCombination) {
					m[newCombination] = struct{}{}
					q.Push(&newCombination)
				}

			}
		}
		count++
	}

	return -1
}

func inDeadends(dedadends [][4]int, comb [4]int) bool {
	for i := range dedadends {
		if dedadends[i] == comb {
			return true
		}
	}
	return false
}
func splitDeadends(deadends []string) [][4]int {
	ddns := make([][4]int, 0, len(deadends))
	for _, num := range deadends {
		newDeadend := split(num)
		ddns = append(ddns, newDeadend)
	}
	return ddns
}

func split(num string) [4]int {
	strs := strings.Split(num, "")
	newDeadend := [4]int{}
	for i := range strs {
		val, _ := strconv.Atoi(strs[i])
		newDeadend[i] = val
	}
	return newDeadend
}

func splitInt(n int) [4]int {
	slc := []int{}
	for n > 0 {
		slc = append(slc, n%10)
		n /= 10
	}

	result := [4]int{}
	for i := range slc {
		result[i] = slc[len(slc)-1-i]
	}

	return result
}

type queue struct {
	arr []*[4]int
	l   *list.List
}

func (q *queue) Len() int { return len(q.arr) }
func (q *queue) Push(el *[4]int) {
	q.arr = append(q.arr, el)
}
func (q *queue) Pop() *[4]int {
	if q.Len() == 0 {
		return nil
	}
	el := q.arr[0]
	q.arr = q.arr[1:]
	return el
}
