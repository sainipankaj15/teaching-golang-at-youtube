package main

import "fmt"

func main() {

	var arr []int
	arr = append(arr, 101)
	arr = append(arr, 202)
	arr = append(arr, 303)
	arr = append(arr, 404)
	arr = append(arr, 505)

	for i := 0; i < len(arr); i++ {
		fmt.Printf("Element at index %d: %d\n", i, arr[i])
	}

	for index, value := range arr {
		fmt.Printf("Element at index %d: %d\n", index, value)
	}
}

