package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println(makeAbbr("abbbvvv"))
	vw := Constructor([]string{"deer", "door", "cake", "card"})

	fmt.Println(vw.IsUnique("dear"))
	fmt.Println(vw.IsUnique("cart"))
	fmt.Println(vw.IsUnique("cane"))
	fmt.Println(vw.IsUnique("make"))
	fmt.Println(vw.IsUnique("cake"))

	vw1 := Constructor([]string{"a", "a"})
	fmt.Println(vw1.IsUnique("a"))
}

type ValidWordAbbr struct {
	m map[string][]string
}

func Constructor(dictionary []string) ValidWordAbbr {
	m := (make(map[string][]string))
	for _, v := range dictionary {
		m[makeAbbr(v)] = append(m[makeAbbr(v)], v)
	}

	return ValidWordAbbr{m: m}
}

func makeAbbr(word string) string {
	if len(word) <= 2 {
		return word
	}
	str1 := string(word[0])
	str2 := strconv.Itoa(len(word[1 : len(word)-1]))
	str3 := string(word[len(word)-1])
	return str1 + str2 + str3
}

func (this *ValidWordAbbr) IsUnique(word string) bool {
	abr := makeAbbr(word)
	for k, v := range this.m {
		if k == abr {
			for _, vv := range v {
				if vv != word {
					return false
				}
			}

			return true
		}
	}
	return true
}
