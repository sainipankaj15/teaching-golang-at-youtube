package main

import "fmt"

func main() {
	fmt.Println("hello world from youtube")

	mp := make(map[string]float64)

	mp["Nifty"] = 25000.1
	mp["25000CE"] = 125.678
	mp["25000PE"] = 98.45
	
	fmt.Println("Map:", mp)

}
