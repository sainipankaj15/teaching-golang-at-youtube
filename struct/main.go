package main

import "fmt"

type Person struct {
	Name string
	Age  int
}

func main() {
	// fmt.Println("Hi")

	p1 := Person{}
	p1.Name = "John"
	p1.Age = 25

	p2 := Person{Name: "Alice", Age: 30}
	p2.Name = "Pankaj"
	fmt.Println(p1)
	fmt.Println(p2)

}
