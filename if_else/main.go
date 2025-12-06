package main

import "fmt"

func main() {

	ltp := 35

	if ltp > 40 {
		fmt.Println("Stock hits the target price, time to sell!")
	} else if ltp < 20 {
		fmt.Println("Stock is below the buy price, time to buy!")
	} 
}
