package main

import "fmt"

func main() {
	in := []rune{'a', 'a', 'b', ' ', ' ', ' ', 'f', 'g', 'h', ' '}
	URLifySlice(in)
	fmt.Println(string(in))
}
func URLifySlice(input []rune) {
	inWord := false
	slowPtr := len(input) - 1
	for i := len(input) - 1; i >= 0; i-- {
		if input[i] == rune(' ') {
			if inWord {
				input[slowPtr-2] = rune('%')
				input[slowPtr-1] = rune('2')
				input[slowPtr] = rune('0')
				slowPtr -= 3
			}
		} else {
			if !inWord {
				inWord = true
			}
			input[slowPtr] = input[i]
			slowPtr--
		}
	}
}
func urify(in []rune) []rune {
	j := 0
	lastSeen := false
	for i := range in {
		if in[i] == ' ' {
			if !lastSeen {
				in[j] = '%'
				j++
				in[j] = '2'
				j++
				in[j] = '0'
				lastSeen = true
			}
		} else {
			in[j] = in[i]
			j++
			lastSeen = false
		}
	}
	return in[0:j]
}
