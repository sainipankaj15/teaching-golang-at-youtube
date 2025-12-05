package main

import "fmt"

func main() {

	var arr [5]string

	arr[0] = "Apple"
	arr[1] = "Banana"
	arr[2] = "Grapes"
	arr[3] = "Mango"
	arr[4] = "Orange"

	fmt.Println("Array elements:")
	fmt.Println(arr)

	fmt.Println("Length of array:", len(arr))

	// Iterating over the array
	fmt.Println("Iterating over array elements:")
	for i := 0; i < len(arr); i++ {
		fmt.Println(arr[i])
	}

	var slice []string
	slice = append(slice, "Red")
	slice = append(slice, "Green")
	slice = append(slice, "Blue")

	fmt.Println("Slice elements:")
	fmt.Println(slice)

	fmt.Println("Length of slice:", len(slice))

	// Iterating over the slice
	fmt.Println("Iterating over slice elements:")
	for i := 0; i < len(slice); i++ {
		fmt.Println(slice[i])
	}
}

