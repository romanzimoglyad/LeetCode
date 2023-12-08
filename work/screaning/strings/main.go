package main

func main() {
	s := "test"
	println(s[0]) // ответ кандидата
	var sBytes = []byte(s)
	sBytes[0] = 'R'
	s = string(sBytes)
	println(s)
}

/**
func test1(){
	s := "test"
	println(s[0]) // ответ кандидата
	s[0] = "R"
	println(s) /
}

*/
