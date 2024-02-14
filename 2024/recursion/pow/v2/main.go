package main

import "fmt"

func main() {
	fmt.Println(myPow(2, 2))
	fmt.Println(myPow(0.00001, 2))
}

func myPow(x float64, n int) float64 {
	if n == 0 {
		return 1
	}
	if n < 0 {
		return 1 / myPowPos(x, -n)
	}
	return myPowPos(x, n)
}

func myPowPos(x float64, n int) float64 {
	if n == 0 {
		return 1
	}
	if n%2 == 0 {
		return myPowPos(x*x, n/2)
	}
	return x * myPowPos(x*x, (n-1)/2)
}
