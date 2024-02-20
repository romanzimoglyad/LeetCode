// Golang program to illustrate the
// concept of the Pointer to struct
package main

import (
	"fmt"
	"time"
)

func main() {
	durationLayout := "15:04:05"
	duration, err := time.Parse(durationLayout, "24:22:11")
	fmt.Println(err)
	fmt.Println(duration)
}
