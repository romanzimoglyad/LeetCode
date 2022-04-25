package main

import "fmt"

func main() {
	mychannel := make(chan int)

	quitchannel := make(chan int)

	sum := 0

	go func() {
		for i := 0; i <= 4; i++ {

			sum += <-mychannel
			fmt.Println("read ", i)
		}
		fmt.Println(sum)
		quitchannel <- 0
	}()
	SumOfSquares(mychannel, quitchannel)
}
func SumOfSquares(c, quit chan int) {
	y := 0
	for {
		select {
		case c <- y * y:
			y++
		case <-quit:
			return
		}
	}

}
