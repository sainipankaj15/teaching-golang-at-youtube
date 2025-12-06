package main

import "fmt"

func add(a int, b int) (int, error) {
	return a + b, nil
}

func main() {

	a := 10
	b := 20

	output, err := add(a, b)

	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("Sum:", output)

}
