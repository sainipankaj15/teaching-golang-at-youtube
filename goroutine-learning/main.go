package main

import (
	"fmt"
	"time"
)

func writeOddNumber() {
	for i := 1; i <= 10000; i += 2 {
		fmt.Println(i)
	}
}

func writeEvenNumber() {
	for i := 0; i <= 10000; i += 2 {
		fmt.Println(i)
	}
}

func main() {
	fmt.Println("Hi")
	go writeOddNumber()
	go writeEvenNumber()

	// go func(msg string) {
	// 	fmt.Println(msg)
	// }("Anonymous Function")

	time.Sleep(2 * time.Second)
}
