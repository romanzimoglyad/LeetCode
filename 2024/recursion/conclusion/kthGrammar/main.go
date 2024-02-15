package main

import "fmt"

func main() {
	fmt.Println(kthGrammar(30, 434991989))
}

func kthGrammarV1(n int, k int) int {
	return buildArr(n)[k-1]
}

func buildArr(n int) []int {
	if n == 1 {
		return []int{0}
	}
	prev := buildArr(n - 1)
	newArr := make([]int, 0, len(prev)*2)
	for i := 0; i < len(prev); i++ {
		if prev[i] == 0 {
			newArr = append(newArr, 0)
			newArr = append(newArr, 1)
		}
		if prev[i] == 1 {
			newArr = append(newArr, 1)
			newArr = append(newArr, 0)
		}
	}
	return newArr
}

func kthGrammar(n int, k int) int {
	if n == 1 {
		return 0
	}
	if k%2 == 1 {
		new := kthGrammar(n-1, k/2+1)
		if new == 0 {
			return 0
		} else {
			return 1
		}
	}

	new := kthGrammar(n-1, k/2)
	if new == 0 {
		return 1
	} else {
		return 0
	}

}
