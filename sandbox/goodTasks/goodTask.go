package main

import (
	"fmt"
)

type S struct {
	a, b int
}

func (s *S) String() string {
	return fmt.Sprintf("%s", s)
}

func main() {
	t := "cs"
	switch t {
	case "as", "cs":
		fmt.Println("asd")
	}

}
