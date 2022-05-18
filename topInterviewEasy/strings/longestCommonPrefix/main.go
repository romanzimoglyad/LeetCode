package main

import (
	"fmt"
)

func main() {
	fmt.Println(longestCommonPrefix([]string{"", ""}))
}

func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}
	i := 0
	curPrefix := ""
	for {
		if len(strs[0]) <= i {
			return curPrefix
		}
		tmpChar := strs[0][i]
		for j := range strs {
			if len(strs[j]) <= i {
				return curPrefix
			}
			if tmpChar != strs[j][i] {
				return curPrefix
			}
		}
		curPrefix = curPrefix + string(tmpChar)
		i++
	}

}
