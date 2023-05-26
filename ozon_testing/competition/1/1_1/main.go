package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	var plural1 = []string{"s", "x", "z"}
	var plural2 = []string{"sh", "ch"}
	gl := []string{"a", "e", "i", "o", "u", "y"}
	path, err := os.Getwd()
	if err != nil {
		log.Println(err)
	}
	fmt.Println(path)
	f, err := os.Open("ozon_testing/competition/1/1_1/01")
	defer f.Close()
	if err != nil {
		log.Fatal(err)
	}
	in := bufio.NewReader(f)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var testCount int
	fmt.Fscan(in, &testCount)

	words := make([]string, testCount)
	for i := 0; i < testCount; i++ {
		fmt.Fscan(in, &words[i])
	}

	for _, v := range words {
		lastChar := v[len(v)-1]
		last2Char := v[len(v)-2:]
		previousChar := v[len(v)-2 : len(v)-1]
		if contains(plural1, string(lastChar)) || contains(plural2, string(last2Char)) {
			fmt.Fprintln(out, v+"es")
		} else if string(lastChar) == "y" && !contains(gl, previousChar) {
			fmt.Fprintln(out, v[:len(v)-1]+"ies")
		} else {
			fmt.Fprintln(out, v+"s")
		}
	}

}
func contains(s []string, e string) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}
