package main

import "fmt"

func main() {

	i := getRow(4)
	fmt.Println(i)
}

func getRow(rowIndex int) []int {
	resp := make([]int, 0, rowIndex+1)
	m := make(map[rc]int)
	for i := 0; i <= rowIndex; i++ {
		resp = append(resp, help(rowIndex, i, m))
	}
	return resp

}

type rc struct {
	r int
	c int
}

func help(rowIndex int, colInd int, m map[rc]int) int {
	if rowIndex == 0 || colInd == 0 || rowIndex == colInd {
		return 1
	}
	rc := rc{r: rowIndex, c: colInd}
	if val, ok := m[rc]; ok {
		return val
	}

	tmp := help(rowIndex-1, colInd-1, m) + help(rowIndex-1, colInd, m)
	m[rc] = tmp
	return tmp
}
