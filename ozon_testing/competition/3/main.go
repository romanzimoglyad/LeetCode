package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
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
		firstVal := 0
		direction := 0
		count := 0
		res := []int{}
		for k := 0; k < numberCount; k++ {

			if k+1 == numberCount {
				if direction == 0 {
					res = append(res, numbers[k])
					res = append(res, 0)
					break
				} else {
					res = append(res, firstVal)

					res = append(res, count*direction)
					break
				}
			}
			if numbers[k+1]-numbers[k] == 1 && (direction == 1 && numbers[k+1] != 0 || direction == 0) {
				if count == 0 {

					firstVal = numbers[k]
				}
				count++
				direction = 1
			} else if numbers[k+1]-numbers[k] == -1 && (direction == -1 && numbers[k+1] != -1 || direction == 0) {
				if count == 0 {

					firstVal = numbers[k]
				}
				count++
				direction = -1

			} else {
				if direction != 0 {
					res = append(res, firstVal)

					res = append(res, count*direction)

				} else {
					res = append(res, numbers[k])
					res = append(res, 0)
				}
				firstVal = 0
				direction = 0
				count = 0
			}
		}
		fmt.Fprintln(out, len(res))

		fmt.Fprintln(out, strings.Join(intSliceToString(res), " "))
	}
}

func intSliceToString(intSlice []int) []string {
	res := make([]string, len(intSlice))
	for i := range intSlice {
		res[i] = strconv.Itoa(intSlice[i])
	}
	return res
}
