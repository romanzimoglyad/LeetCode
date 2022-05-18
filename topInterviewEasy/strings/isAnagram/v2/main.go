package main

import "fmt"

func main() {
	fmt.Println(isAnagram("aacc", "ccac"))
}

func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	t1 := 'z'
	t2 := 'a'
	t3 := t1 - t2
	fmt.Println(t3)

	letter := make([]int, 26)
	for _, v := range s {
		letter[v-'a']++
	}
	for _, v := range t {
		if letter[v-'a'] <= 0 {
			return false
		}
		letter[v-'a']--
	}
	// for _,i := range(letter){
	//     if i !=0{
	//         return false
	//     }
	// }
	return true
}
