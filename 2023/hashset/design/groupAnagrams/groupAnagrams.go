package main

import (
	"fmt"
	"sort"
)

func main() {
	res := groupAnagrams([]string{"abc", "cba", "aaaa", "aaaa"})
	fmt.Println(res)
}

func groupAnagrams(strs []string) [][]string {
	m := make(map[string][]string)
	for _, v := range strs {
		sorted := sortString(v)
		m[sorted] = append(m[sorted], v)
	}
	res := make([][]string, 0, len(m))
	for k := range m {
		res = append(res, m[k])
	}
	return res
}

func sortString(str string) string {
	b := []byte(str)
	sort.SliceStable(b, func(i, j int) bool {
		return b[i] < b[j]
	})
	return string(b)
}
