package main

import "fmt"

func main() {
	resp := generateTrees(5)
	fmt.Println(resp)
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func generateTrees(n int) []*TreeNode {
	helpApp := make([]int, 0, n)
	for i := 1; i <= n; i++ {
		helpApp = append(helpApp, i)
	}
	resp := []*TreeNode{}
	for i := 1; i <= n; i++ {

		r := rec(helpApp[:i-1], helpApp[i:], &TreeNode{Val: i})
		resp = append(resp, r...)
	}
	return resp
}

func rec(left []int, right []int, val *TreeNode) []*TreeNode {
	resp := []*TreeNode{}
	if len(left) == 0 && len(right) == 0 {
		return []*TreeNode{val}
	}
	if len(left) == 0 {
		for k, v := range right {
			rightResp := rec(right[:k], right[k+1:], &TreeNode{Val: v})

			for i := range rightResp {
				newVal := *val
				newVal.Right = rightResp[i]
				resp = append(resp, &newVal)
			}
		}
	} else if len(right) == 0 {
		for k, v := range left {
			leftResp := rec(left[:k], left[k+1:], &TreeNode{Val: v})

			for i := range leftResp {
				newVal := *val
				newVal.Left = leftResp[i]
				resp = append(resp, &newVal)
			}
		}
	} else {
		for kl, vl := range left {
			for kr, vr := range right {
				rightResp := rec(right[:kr], right[kr+1:], &TreeNode{Val: vr})
				leftResp := rec(left[:kl], left[kl+1:], &TreeNode{Val: vl})
				for i := range rightResp {
					for k := range leftResp {
						newVal := *val
						newVal.Right = rightResp[i]
						newVal.Left = leftResp[k]
						resp = append(resp, &newVal)
					}
				}
			}
		}
	}
	return resp
}
