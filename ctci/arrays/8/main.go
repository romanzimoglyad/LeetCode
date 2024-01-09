package main

import "fmt"

func main() {
	in := [][]int{{0, 1, 2}, {0, 2, 3}, {1, 2, 6}}
	zero(in)
	fmt.Println(in)
}

func zero(m [][]int) {
	rowM := make(map[int]struct{})
	colM := make(map[int]struct{})
	for i := range m {
		for j := range m[i] {
			if m[i][j] == 0 {
				rowM[i] = struct{}{}
				colM[j] = struct{}{}
			}
		}
	}
	
	for i := range rowM {
		for j := range m[i] {
			m[i][j] = 0
		}
	}

	for j := range colM {
		for i := 0; i < len(m); i++ {
			m[i][j] = 0
		}
	}

}
