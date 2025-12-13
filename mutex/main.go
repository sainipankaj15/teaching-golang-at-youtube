// Online Go compiler to run Golang program online
// Print "Try programiz.pro" message

package main

import (
	"fmt"
	"sync"
)

type Person struct {
	Name string
	Age  int
}

func main() {
	fmt.Println("Try programiz.pro")

	person1 := Person{}

	//   var mu sync.Mutex
	var mu sync.RWMutex

	mu.Lock()
	person1.Name = "Pankaj"
	person1.Age = 19
	mu.Unlock()

	mu.RLock()
	p2 := person1
	p2.Age = 20
	mu.RUnlock()

	fmt.Println(person1)
	fmt.Println(p2)
}
