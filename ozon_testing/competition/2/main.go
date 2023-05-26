package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var testCount int
	fmt.Fscan(in, &testCount)

	for i := 0; i < testCount; i++ {
		numberCount := 0
		fmt.Fscan(in, &numberCount)
		numbers := make([]int, numberCount)
		for i := 0; i < numberCount; i++ {
			fmt.Fscan(in, &numbers[i])
		}

		allChilds := map[int]struct{}{}
		var v int
		var elems []elem
		for v != numberCount {
			val := numbers[v]
			elem := elem{edge: val}
			count := numbers[v+1]
			for i := 0; i < int(count); i++ {
				allChilds[numbers[v+2+i]] = struct{}{}
			}
			elems = append(elems, elem)
			v = v + count + 2

		}
		for _, v := range elems {
			val := v.edge
			if _, ok := allChilds[val]; !ok {
				fmt.Fprintln(out, val)
				break
			}
		}
	}
}

type elem struct {
	edge int
}
