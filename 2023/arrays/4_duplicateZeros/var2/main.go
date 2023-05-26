package main

import "fmt"

func main() {
	testArr := []int{0, 0, 0}
	duplicateZeros(testArr)
	fmt.Println(testArr)
}
func duplicateZeros(arr []int) {
	duplZeros := 0
	l := len(arr) - 1
	for i := 0; i <= l-duplZeros; i++ {
		if arr[i] == 0 {
			if i == l-duplZeros {
				arr[l] = 0
				l -= 1
				fmt.Println("asd")
				break
			}
			duplZeros++
		}
	}

	for i := l - duplZeros; i >= 0; i-- {
		if arr[i] != 0 {
			arr[i+duplZeros] = arr[i]

		} else {
			arr[i+duplZeros] = 0
			duplZeros--
			arr[i+duplZeros] = 0
		}

	}

}
