package oneWay

import (
	"testing"
)

func TestAreOneEditAway(t *testing.T) {
	cases := []struct {
		number   int
		input1   string
		input2   string
		expected bool
	}{
		{1, "abcd", "abcd", true},
		{2, "abcd", "abcc", true},
		{3, "abcd", "accc", false},
		{4, "abcd", "abcde", true},
		{5, "abcd", "abcdef", false},
		{6, "abcde", "abcd", true},
		{7, "abcdef", "abcd", false},
		{8, " ", "", true},
		{9, "", " ", true},
		{10, "", "", true},
	}
	for _, c := range cases {
		actual := oneWay(c.input1, c.input2)
		if actual != c.expected {
			t.Fatalf("Number:%v,Inputs %s, %s. Expected: %v, actual: %v\n",
				c.number, c.input1, c.input2, c.expected, actual)
		}
	}
}
