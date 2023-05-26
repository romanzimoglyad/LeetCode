package main

import (
	"bufio"
	"fmt"
	"os"
)

const (
	ship1 = 1
	ship2 = 2
	ship3 = 3
	ship4 = 4
	yes   = "YES"
	no    = "NO"
)

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var testCount int
	fmt.Fscan(in, &testCount)

	for i := 0; i < testCount; i++ {
		var number1, number2, number3, number4 int
		a := make([]uint32, 10)
		for i := 0; i < 10; i++ {
			fmt.Fscan(in, &a[i])
		}
		for _, val := range a {
			if val == ship1 {
				number1++
			}
			if val == ship2 {
				number2++
			}
			if val == ship3 {
				number3++
			}
			if val == ship4 {
				number4++
			}
		}
		fmt.Fprintln(out, checks(number1, number2, number3, number4))
	}
}

func checks(number1, number2, number3, number4 int) string {
	if number1 != ship4 {
		return no
	}
	if number2 != ship3 {
		return no
	}
	if number3 != ship2 {
		return no
	}
	if number4 != ship1 {
		return no
	}
	return yes
}
