package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
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
	path, err := os.Getwd()
	if err != nil {
		log.Println(err)
	}
	fmt.Println(path)
	f, err := os.Open("ozon_testing/7_build_prder/7_2/file.txt")
	defer f.Close()
	if err != nil {
		log.Fatal(err)
	}
	in := bufio.NewReader(f)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var testCount int
	fmt.Fscan(in, &testCount)
	for i := 0; i < testCount; i++ {
		result := "x"
		var windowsNumber, pationsNumber int
		fmt.Fscan(in, &windowsNumber, &pationsNumber)
		beginning := make([]int, pationsNumber)
		for i := 0; i < pationsNumber; i++ {
			fmt.Fscan(in, &beginning[i])
		}
		beginingUnits := make([]Unit, pationsNumber)
		for i := range beginning {
			beginingUnits[i] = Unit{
				value:       beginning[i],
				oldPosition: i,
			}
		}

		sort.Slice(beginingUnits, func(i, j int) bool {
			return beginingUnits[i].value < beginingUnits[j].value
		})
		if isUnique(beginingUnits) {
			result = convertResponse(beginingUnits, pationsNumber)
			goto end
		}

		for i := 0; i < pationsNumber; i++ {
			if beginingUnits[i].value-1 > 0 && (i-1 < 0 || beginingUnits[i].value-1 > beginingUnits[i-1].value) {
				beginingUnits[i].value--
				beginingUnits[i].move = -1
			} else if i-1 < 0 || beginingUnits[i].value != beginingUnits[i-1].value {
				beginingUnits[i].move = 0
			} else if beginingUnits[i].value+1 <= windowsNumber {
				beginingUnits[i].value++
				beginingUnits[i].move = 1
			} else {
				goto end
			}
		}
		if isUnique(beginingUnits) {
			result = convertResponse(beginingUnits, pationsNumber)
			goto end
		}
	end:
		fmt.Fprintln(out, result)
	}
}
func convertResponse(in []Unit, n int) string {
	res := make([]int, n)
	for i := range in {
		res[in[i].oldPosition] = in[i].move
	}
	resp := strings.Builder{}
	for i := range res {
		switch res[i] {
		case 0:
			resp.WriteString("0")
		case 1:
			resp.WriteString("+")
		case -1:
			resp.WriteString("-")
		}
	}
	return resp.String()
}
func isUnique(slice []Unit) bool {
	seen := make(map[int]bool)
	for _, elem := range slice {
		if seen[elem.value] {
			return false
		}
		seen[elem.value] = true
	}
	return true
}
func findOldPos(beginning []int, val int) int {
	for i := range beginning {
		if beginning[i] == val {
			return i
		}
	}
	return -1
}

type Unit struct {
	value       int
	oldPosition int
	move        int
}
