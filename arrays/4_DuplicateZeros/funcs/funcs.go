package funcs

func main() {

}
func DuplicateZeros1(arr []int) {
	res := make([]int, 0, len(arr))
	for i := 0; i < len(arr); i++ {
		if arr[i] != 0 {
			res = append(res, arr[i])
		} else {
			res = append(res, arr[i])

			res = append(res, 0)

		}
	}
	copy(arr, res)

}
func DuplicateZeros2(arr []int) {
	res := make([]int, 0, len(arr)*2)
	for i := 0; i < len(arr); i++ {
		if arr[i] != 0 {
			res = append(res, arr[i])
		} else {
			res = append(res, arr[i])

			res = append(res, 0)

		}
	}
	copy(arr, res)

}
