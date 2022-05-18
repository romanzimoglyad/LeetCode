package main

import "fmt"

func main() {
	fmt.Println(countPrimes(10))
}

func countPrimes(n int) int {
	count := 0
	isPrime := make([]bool, n)
	for i := 2; i < n; i++ {
		isPrime[i] = true
	}
	for i := 2; i*i < n; i++ {
		if !isPrime[i] {
			continue
		}
		for j := i * i; j < n; j += i {
			isPrime[j] = false
		}

	}
	for i := range isPrime {
		if isPrime[i] {
			count++
		}
	}
	return count
}
