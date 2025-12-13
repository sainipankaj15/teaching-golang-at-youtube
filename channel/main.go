package main

import (
	"fmt"
)

// func willPushNumber(ch chan int) {
// 	ch <- 100
// }

func main() {

	ch := make(chan int, 2)
	ch <- 100
	ch <- 200

	num := <-ch
	fmt.Printf("Received number: %d\n", num)
	num = <-ch
	fmt.Printf("Received number: %d\n", num)

	ch <- 300
	num = <-ch
	fmt.Printf("Received number: %d\n", num)
}
