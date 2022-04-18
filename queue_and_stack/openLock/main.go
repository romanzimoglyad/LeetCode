package main

import (
	"fmt"
	"strconv"
)

func main() {
	deadends := []string{"0201", "0101", "0102", "1212", "2002"}
	target := "0202"

	fmt.Println(openLock(deadends, target))

}

var directions = []int{1, -1}

func openLock(deadends []string, target string) int {

	if contains(deadends, "0000") {
		return -1
	}
	if target == "0000" {
		return 0
	}
	queue := newQueue()
	m := map[string]struct{}{}
	queue.add("0000")
	queue.add("")
	m["0000"] = struct{}{}
	step := 0
	for queue.len() > 0 {
		number := queue.pop()
		if number == "" {
			step++
			queue.add("")
		}
		for i, val := range number {
			for _, d := range directions {
				newNumber := getNewNumber(val, i, d, number)
				if _, ok := m[newNumber]; ok || contains(deadends, newNumber) || newNumber == "" {
					continue
				}
				if newNumber == target {
					return step + 1
				}
				m[newNumber] = struct{}{}
				queue.add(newNumber)
			}
		}

	}
	return -1
}

func getNewNumber(val int32, i int, d int, number string) string {
	n, _ := strconv.Atoi(string(val))
	newVal := n + d
	if newVal == -1 {
		newVal = 9
	}
	newNumber := number[:i] + strconv.Itoa(newVal%10) + number[i+1:]
	in, _ := strconv.Atoi(newNumber)
	if in > 9999 {
		return ""
	}
	return newNumber
}

type stack struct {
	nodes  []string
	length int
}

func newQueue() stack {
	return stack{
		nodes:  []string{},
		length: 0,
	}
}
func (q *stack) len() int {
	return q.length

}
func (q *stack) add(elem string) {
	q.nodes = append(q.nodes, elem)
	q.length++
}
func (q *stack) pop() string {
	if q.length == 0 {
		return ""
	}
	res := q.nodes[0]
	q.nodes = q.nodes[1:len(q.nodes)]
	q.length--
	return res
}

func contains(slice []string, target string) bool {
	for i := range slice {
		if slice[i] == target {
			return true
		}
	}
	return false
}
