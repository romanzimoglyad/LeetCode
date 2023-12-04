package main

import "fmt"

func main() {
	nums := []int{1, 1, 1, 1, 1}
	res := findTargetSumWaysV2(nums, 3)
	fmt.Println(res)
}

func findTargetSumWaysV2(nums []int, target int) int {
	length := len(nums)
	result := 0
	rec(nums, nums[0], 0, target, length-1, &result)
	rec(nums, -nums[0], 0, target, length-1, &result)
	return result
}

func rec(nums []int, cur int, ind int, res int, length int, result *int) {
	if ind == length {
		if cur == res {
			*result++
		}
		return
	}
	rec(nums, cur+nums[ind+1], ind+1, res, length, result)
	rec(nums, cur-nums[ind+1], ind+1, res, length, result)
}

func findTargetSumWays(nums []int, target int) int {
	length := len(nums)
	result := 0
	stack := Constructor()
	stack.Push(&elem{
		value: nums[0],
		pos:   0,
	})
	stack.Push(&elem{
		value: -nums[0],
		pos:   0,
	})
	for stack.Top() != nil {
		topElem := stack.Top()
		stack.Pop()
		if topElem.pos == length-1 {
			if topElem.value == target {
				result++
			}
			continue
		}
		stack.Push(&elem{
			value: topElem.value + nums[topElem.pos+1],
			pos:   topElem.pos + 1,
		})
		stack.Push(&elem{
			value: topElem.value - nums[topElem.pos+1],
			pos:   topElem.pos + 1,
		})
	}
	return result
}

type elem struct {
	value int
	pos   int
}

type stack struct {
	arr    []*elem
	length int
}

func Constructor() stack {
	return stack{arr: make([]*elem, 0)}
}

func (this *stack) Push(val *elem) {
	this.arr = append(this.arr, val)

	this.length++
}

func (this *stack) Pop() {
	this.arr = this.arr[:this.length-1]
	this.length--
}

func (this *stack) Top() *elem {
	if this.length == 0 {
		return nil
	}
	return this.arr[this.length-1]
}
