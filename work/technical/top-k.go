package technical

func kMostFreq(arr []int, k int) []int {
	freqMap := make(map[int]int)

	for i := 0; i < len(arr); i++ {
		freqMap[arr[i]]++
	}
	nums := make([][]int, len(arr)+1)

	for k, v := range freqMap {
		nums[v] = append(nums[v], k)
	}
	res := make([]int, 0, k)

	i := len(arr)
	for k > len(res) && i > 0 {
		v := nums[i]
		if v != nil {
			res := append(res, v...)
		}
		i--
	}

	return res
}
