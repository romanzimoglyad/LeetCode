package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Printed struct {
	separatedPages []int
	diapasons      []Diapason
}
type Diapason struct {
	start int
	end   int
}

var directions = []int{1, 0, -1}

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var testCount int
	fmt.Fscan(in, &testCount)

	for i := 0; i < testCount; i++ {
		var windowsNumber, pationsNumber int
		fmt.Fscan(in, &windowsNumber, &pationsNumber)
		times := make([]int, pationsNumber)
		for i := 0; i < pationsNumber; i++ {
			fmt.Fscan(in, &times[i])
		}
		stack := newStack()
		try := TryInfo{
			indexes: make([]int, pationsNumber),
			slice:   times,
			moves:   make([]int, pationsNumber),
		}
		stack.add(&try)
		result := "x"
		for stack.len() > 0 {
			el := stack.pop()
			if el == nil {

				break
			}
			for elKey, _ := range el.slice {
				for _, dirVal := range directions {
					newTry := getNewTry(*el, elKey, dirVal, windowsNumber)
					if newTry == nil {
						break
					}
					if isUnique(newTry.slice) {
						result = convertResponse(newTry.moves)
						goto end
					}
					stack.add(newTry)

				}
			}

		}
	end:
		fmt.Fprintln(out, result)
	}
}

func convertResponse(in []int) string {
	res := strings.Builder{}
	for i := range in {
		switch in[i] {
		case 0:
			res.WriteString("0")
		case 1:
			res.WriteString("+")
		case -1:
			res.WriteString("-")
		}
	}
	return res.String()
}

func isUnique(slice []int) bool {
	seen := make(map[int]bool)
	for _, elem := range slice {
		if seen[elem] {
			return false
		}
		seen[elem] = true
	}
	return true
}

func getNewTry(old TryInfo, key, dir, windowsNumber int) *TryInfo {
	if old.indexes[key] == 1 {
		return nil
	}
	new := make([]int, len(old.slice))
	copy(new, old.slice)
	newMoves := make([]int, len(old.moves))
	copy(newMoves, old.moves)
	newIndexes := make([]int, len(old.moves))
	copy(newIndexes, old.indexes)
	newVal := new[key] + dir
	if newVal > windowsNumber || newVal == 0 {
		return nil
	}
	newTryInfo := old
	new[key] = newVal
	newMoves[key] = dir
	newIndexes[key] = 1
	newTryInfo.slice = new
	newTryInfo.moves = newMoves
	newTryInfo.indexes = newIndexes
	return &newTryInfo
}

type stack struct {
	nodes  []*TryInfo
	length int
}

type TryInfo struct {
	indexes []int
	slice   []int
	moves   []int
}

func newStack() stack {
	return stack{
		nodes:  []*TryInfo{},
		length: 0,
	}
}
func (q *stack) len() int {
	return q.length

}
func (q *stack) add(elem *TryInfo) {
	q.nodes = append(q.nodes, elem)
	q.length++
}
func (q *stack) pop() *TryInfo {
	if q.length == 0 {
		return nil
	}
	res := q.nodes[0]
	q.nodes = q.nodes[1:len(q.nodes)]
	q.length--
	return res
}

func contains(slice []int, target int) bool {
	for i := range slice {
		if slice[i] == target {
			return true
		}
	}
	return false
}
