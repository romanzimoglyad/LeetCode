package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	path, err := os.Getwd()
	if err != nil {
		log.Println(err)
	}
	fmt.Println(path)
	f, err := os.Open("ozon_testing/5/7_1/file.txt")
	defer f.Close()
	if err != nil {
		log.Fatal(err)
	}
	in := bufio.NewReader(f)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var testCount int
	fmt.Fscan(in, &testCount)

	for i := 0; i < testCount; i++ {
		var count int
		fmt.Fscan(in, &count)

		signs := make([]string, count)

		for i := 0; i < count; i++ {
			fmt.Fscan(in, &signs[i])
		}
		var registered []string
		for _, newSign := range signs {
			if len(registered) == 0 {
				registered = append(registered, newSign)
				continue
			}
			similuar := false
			for _, oldSign := range registered {
				if newSign == oldSign {
					similuar = true
					break
				}
				if signsSimilar(oldSign, newSign) {
					similuar = true
					break
				}

			}
			if !similuar {
				registered = append(registered, newSign)
			}
		}
		fmt.Fprintln(out, len(registered))
	}

}

func signsSimilar(oldSign, newSign string) bool {
	if (len(oldSign) == 1 || len(newSign) == 1) && oldSign != newSign {
		return false
	}
	var previousLetter uint8
	var previousLetterCount int
	for i, j := 0, 0; i < len(oldSign) || j < len(newSign); {
		if j == len(newSign) {
			for k := i; k < len(oldSign); k++ {
				if oldSign[k] != previousLetter || (oldSign[k] == previousLetter && previousLetterCount < 2) {
					return false
				}
			}
			return true
		}
		if i == len(oldSign) {
			for k := j; k < len(newSign); k++ {
				n := newSign[k]
				if n != previousLetter || (n == previousLetter && previousLetterCount < 2) {
					return false
				}
			}
			return true
		}
		oldLetter := oldSign[i]
		newLetter := newSign[j]
		if oldLetter == newLetter {
			i++
			j++
			if oldLetter == previousLetter {
				previousLetterCount++
			} else {
				previousLetter = oldLetter
				previousLetterCount = 1
			}
		} else {
			if previousLetter == oldLetter && previousLetterCount >= 2 {
				i++
			} else if previousLetter == newLetter && previousLetterCount >= 2 {
				j++
			} else {
				return false
			}
		}
	}
	return true
}
