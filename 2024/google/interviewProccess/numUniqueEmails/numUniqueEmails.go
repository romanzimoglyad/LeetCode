package main

import (
	"fmt"
	"strings"
)

func main() {
	test := []string{"a.as.das+as.dasd@leetcode.com", "b@leetcode.com", "c@leetcode.com"}
	resp := numUniqueEmails(test)
	fmt.Println(resp)
}

const dot = '.'
const plus = '+'
const dog = '@'

func numUniqueEmails(emails []string) int {
	m := make(map[string]struct{})
	for _, email := range emails {
		dogIndex := strings.IndexRune(email, dog)
		if dogIndex == -1 {
			continue
		}
		localName := email[:dogIndex]
		domain := email[dogIndex+1:]
		indexPlus := strings.IndexRune(localName, plus)
		if indexPlus != -1 {
			localName = localName[0:indexPlus]
		}
		localName = strings.ReplaceAll(localName, ".", "")
		m[localName+"@"+domain] = struct{}{}
	}
	return len(m)
}
