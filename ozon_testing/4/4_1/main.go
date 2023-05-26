package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	path, err := os.Getwd()
	if err != nil {
		log.Println(err)
	}
	fmt.Println(path)
	f, err := os.Open("ozon_testing/4/4_1/file.txt")
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
		guys := make([]int, count)
		for i := 0; i < count; i++ {
			fmt.Fscan(in, &guys[i])
		}
		newGuys := make([]int, count)
		copy(newGuys, guys)
		sort.Ints(newGuys)
		prev := 0
		m := make(map[int]int)
		var place = 0
		var ec = 0
		first := true
		for _, guy := range newGuys {
			if guy-prev > 1 || first {
				first = false
				place++
				m[guy] = place + ec
				place = place + ec
				ec = 0

			} else {
				ec++
				m[guy] = place
			}
			prev = guy
		}
		resp := make([]string, 0, count)
		for _, guy := range guys {
			v, _ := m[guy]
			resp = append(resp, strconv.Itoa(v))
		}
		fmt.Fprintln(out, strings.Join(resp, " "))
	}
}
