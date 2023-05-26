package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strings"
)

var pattern1 = regexp.MustCompile(`[A-Za-z]\d{2}[A-Za-z]{2}`)
var pattern2 = regexp.MustCompile(`[A-Za-z]\d[A-Za-z]{2}`)

const (
	length1 = 5
	length2 = 4
)

func main() {
	path, err := os.Getwd()
	if err != nil {
		log.Println(err)
	}
	fmt.Println(path)
	f, err := os.Open("ozon_testing/3/file.txt")
	defer f.Close()
	if err != nil {
		log.Fatal(err)
	}
	in := bufio.NewReader(f)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var testCount int
	fmt.Fscan(in, &testCount)
next:
	for i := 0; i < testCount; i++ {
		var line string
		fmt.Fscan(in, &line)

		s := make([]string, 0, 50)
		for len(line) >= length2 {
			if pattern2.MatchString(line[0:length2]) {
				s = append(s, line[0:length2])

				line = line[length2:]
			} else if len(line) >= length1 && pattern1.MatchString(line[0:length1]) {
				s = append(s, line[0:length1])
				line = line[length1:]
			} else {
				fmt.Fprintln(out, "-")
				goto next
			}
		}
		if len(line) != 0 {
			fmt.Fprintln(out, "-")
		} else {
			fmt.Fprintln(out, strings.Join(s, " "))
		}
	}
}
