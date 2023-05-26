package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

var left uint8 = 'L'
var right uint8 = 'R'
var up uint8 = 'U'
var down uint8 = 'D'
var home uint8 = 'B'
var end uint8 = 'E'
var enter uint8 = 'N'

func main() {

	path, err := os.Getwd()
	if err != nil {
		log.Println(err)
	}
	fmt.Println(path)
	f, err := os.Open("ozon_testing/competition/4/2_2/01")
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

		var line string
		fmt.Fscan(in, &line)

		//m := map[int][]string{}
		cond := condition{
			//linesToText: m,
			sliceLinesToText: [][]string{}}
		cond.sliceLinesToText = append(cond.sliceLinesToText, []string{})
		for i := 0; i < len(line); i++ {
			switch line[i] {
			case left:

				if cond.cursorPos != 0 {
					cond.cursorPos--
				}
			case right:
				if len(cond.sliceLinesToText[cond.cursorLine]) != cond.cursorPos {
					cond.cursorPos++
				}
			case up:

				if cond.cursorLine >= 1 {
					cond.cursorLine--
					if len(cond.sliceLinesToText[cond.cursorLine]) <= cond.cursorPos {
						cond.cursorPos = len(cond.sliceLinesToText[cond.cursorLine])
					}
				}
			case down:

				if len(cond.sliceLinesToText)-cond.cursorLine-1 != 0 {
					cond.cursorLine++
					if len(cond.sliceLinesToText[cond.cursorLine]) <= cond.cursorPos {
						cond.cursorPos = len(cond.sliceLinesToText[cond.cursorLine])
					}
				}
			case home:
				cond.cursorPos = 0
			case end:

				cond.cursorPos = len(cond.sliceLinesToText[cond.cursorLine])
			case enter:
				cond.enterNumbers++
				cond.totalLines++

				if len(cond.sliceLinesToText[cond.cursorLine]) != cond.cursorPos {
					previousToAdd := cond.sliceLinesToText[cond.cursorLine][cond.cursorPos:]
					previousToStay := cond.sliceLinesToText[cond.cursorLine][:cond.cursorPos]
					cond.sliceLinesToText[cond.cursorLine] = previousToStay
					cond.cursorLine++
					if len(cond.sliceLinesToText) <= cond.cursorLine {
						cond.sliceLinesToText = append(cond.sliceLinesToText, []string{})
						cond.sliceLinesToText[cond.cursorLine] = []string{}

					} else {
						cond.sliceLinesToText = append(cond.sliceLinesToText[:cond.cursorLine+1], cond.sliceLinesToText[cond.cursorLine:]...)
						cond.sliceLinesToText[cond.cursorLine] = []string{}

					}

					newVal := cond.sliceLinesToText[cond.cursorLine]
					newVal = append(newVal, previousToAdd...)
					cond.sliceLinesToText[cond.cursorLine] = newVal
					cond.cursorPos = 0
					//	val := cond.sliceLinesToText[cond.cursorLine]

				} else {
					cond.cursorLine++
					cond.cursorPos = 0
					cond.sliceLinesToText = append(cond.sliceLinesToText, []string{})
				}
			default:

				//curLine := cond.linesToText[cond.cursorLine]

				curLine := cond.sliceLinesToText[cond.cursorLine]
				pos := cond.cursorPos
				if len(curLine) == pos {
					//curLine = append(curLine, string(line[i]))
					curLine = append(curLine, string(line[i]))
				} else {
					curLine = append(curLine[:pos+1], curLine[pos:]...)
					curLine[pos] = string(line[i])
				}
				cond.cursorPos++
				cond.sliceLinesToText[cond.cursorLine] = curLine
			}
		}

		if len(cond.sliceLinesToText) == 0 {
			fmt.Fprintln(out, "")
		} else {
			for _, v := range cond.sliceLinesToText {
				if len(v) != 0 {
					fmt.Fprintln(out, strings.Join(v, ""))
				} else {
					fmt.Fprintln(out, "")
				}
			}
		}

		fmt.Fprintln(out, "-")
	}
}

type condition struct {
	totalLines int
	//	linesToText      map[int][]string
	sliceLinesToText [][]string
	texts            [][]string
	cursorLine       int
	cursorPos        int
	enterNumbers     int
}
