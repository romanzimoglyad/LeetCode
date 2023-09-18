package main

import (
	"fmt"
	"strconv"
)

func main() {
	res := addBinary("1111", "1111")
	fmt.Println(res)
}

func addBinary(a string, b string) string {
	aLen := len(a)
	bLen := len(b)
	//minLen := minLen(aLen, bLen)
	aBytes := []byte(a)
	bBytes := []byte(b)
	var inMemory int

	var response []int
	aInt := aLen - 1
	bInt := bLen - 1
	for aInt >= 0 || bInt >= 0 {
		aVal := 0
		if aInt >= 0 {
			aVal, _ = strconv.Atoi(string(aBytes[aInt]))

		}
		bVal := 0
		if bInt >= 0 {
			bVal, _ = strconv.Atoi(string(bBytes[bInt]))
		}
		res := aVal + bVal + inMemory
		if res == 3 {
			response = append(response, 1)
			inMemory = 1
		} else if res == 2 {
			response = append(response, 0)
			inMemory = 1
		} else if res == 1 {
			response = append(response, 1)
			inMemory = 0
		} else {
			response = append(response, 0)
			inMemory = 0
		}
		aInt--
		bInt--
	}
	if inMemory == 1 {
		response = append(response, 1)
	}
	for i, j := 0, len(response)-1; i < j; i, j = i+1, j-1 {
		response[i], response[j] = response[j], response[i]
	}
	resp := ""
	for i := range response {
		resp += strconv.Itoa(response[i])
	}
	return resp

}
