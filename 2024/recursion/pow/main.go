package main

import "fmt"

func main() {
	fmt.Println(myPow(2, 2))
	fmt.Println(myPow(0.00001, 2))
}

func myPow(x float64, n int) float64 {

	if n < 0 {
		return myPowPos(1/x, 1/x, -n)
	}
	return myPowPos(x, x, n)
}

func myPowPos(answer, x float64, n int) float64 {
	if n == 1 {
		return answer
	}
	return myPowPos(answer*x, x, n-1)
}
