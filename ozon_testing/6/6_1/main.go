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

type Printed struct {
	separatedPages []int
	diapasons      []Diapason
}
type Diapason struct {
	start int
	end   int
}

func main() {
	path, err := os.Getwd()
	if err != nil {
		log.Println(err)
	}
	fmt.Println(path)
	f, err := os.Open("ozon_testing/6/7_1/file.txt")
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
		var pagesCount int
		fmt.Fscan(in, &pagesCount)
		var format string
		fmt.Fscan(in, &format)
		alreadyPrinted := parseFormat(format)
		needToPrint := alreadyPrinted.getAllUnPrintedPages(pagesCount)
		fmt.Fprintln(out, convert(needToPrint))
	}
}

func convert(ints []int) string {
	var result strings.Builder
	start := -1
	end := -1
	for i := 0; i <= len(ints)-1; i++ {
		if i+1 == len(ints) {
			if start == -1 {
				result.WriteString("," + strconv.Itoa(ints[i]))
				break
			} else if ints[i] == end {
				result.WriteString("," + strconv.Itoa(start) + "-" + strconv.Itoa(ints[i]))
				break
			}
		}
		if ints[i] == ints[i+1]-1 {
			if start == -1 {
				start = ints[i]
			}
			end = ints[i+1]
		} else {
			if start != -1 && end != -1 {
				result.WriteString("," + strconv.Itoa(start) + "-" + strconv.Itoa(end))
				start, end = -1, -1
			} else {
				result.WriteString("," + strconv.Itoa(ints[i]))
			}
		}

	}
	return result.String()[1:]
}

func (p *Printed) getAllUnPrintedPages(n int) []int {
	for _, v := range p.diapasons {
		for i := v.start; i <= v.end; i++ {
			p.separatedPages = append(p.separatedPages, i)
		}
	}
	sort.Ints(p.separatedPages)
	res := []int{}
	for i := 1; i <= n; i++ {
		if !containsInSlice(p.separatedPages, i) {
			res = append(res, i)
		}
	}
	return res
}

func containsInSlice(s []int, e int) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

func makeNewRequest(printed Printed) string {
	/*for k, v := range printed.separatedPages {

	}*/
	return ""
}

func parseFormat(format string) Printed {
	printed := Printed{}
	strs := strings.Split(format, ",")
	for _, v := range strs {
		if val, err := strconv.Atoi(v); err == nil {
			printed.separatedPages = append(printed.separatedPages, val)
		} else {
			start, err := strconv.Atoi(v[0:strings.Index(v, "-")])
			if err != nil {
				panic(err)
			}
			end, err := strconv.Atoi(v[strings.Index(v, "-")+1:])
			if err != nil {
				panic(err)
			}
			printed.diapasons = append(printed.diapasons, Diapason{start: start, end: end})
		}
	}

	return printed
}
