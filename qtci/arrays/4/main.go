package oneWay

import (
	"math"
	"strings"
)

func oneWay(in1, in2 string) bool {
	l1 := len(in1)
	l2 := len(in2)
	in1Arr := strings.Split(in1, "")
	in2Arr := strings.Split(in2, "")
	if l1 == l2 {
		return checkSameLen(in1Arr, in2Arr)
	}
	if math.Abs(float64(l1-l2)) > 1 {
		return false
	}
	return checkDiffLen(in1Arr, l1, in2Arr, l2)
}

func checkDiffLen(in1 []string, l1 int, in2 []string, l2 int) bool {

	if l1 > l2 {
		return check(in1, l1, in2, l2)
	}
	return check(in2, l2, in1, l1)
}

func check(in1 []string, l1 int, in2 []string, l2 int) bool {
	action := false
	j := 0
	for i := 0; i < l1; i++ {
		if j == l2 || in1[i] != in2[j] {
			if action {
				return false
			}
			action = true
		} else {
			j++
		}
	}
	return true
}

func checkSameLen(in1, in2 []string) bool {
	action := false
	for i := range in1 {
		if in1[i] != in2[i] {
			if action {
				return false
			}
			action = true
		}
	}
	return true
}
