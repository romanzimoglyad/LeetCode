package funcs

func SortedSquares1(nums []int) []int {
	tempRes := []int{}

	for i := range nums {
		newVar := nums[i] * nums[i]
		tempRes = append(tempRes, newVar)
	}

	for i := 0; i < len(tempRes); i++ {
		for j := len(tempRes) - 1; j > i; j-- {
			if tempRes[j] < tempRes[i] {
				tmp := tempRes[i]
				tempRes[i] = tempRes[j]
				tempRes[j] = tmp
			}
		}
	}
	return tempRes
}
func SortedSquares2(nums []int) []int {
	tempRes := make([]int, len(nums))

	left := 0
	right := len(nums) - 1
	i := len(nums)
	for left <= right {
		i--
		newLeft := nums[left] * nums[left]
		newRight := nums[right] * nums[right]
		if newRight > newLeft {
			tempRes[i] = newRight
			right--
		} else {
			tempRes[i] = newLeft
			left++
		}
	}
	return tempRes
}
