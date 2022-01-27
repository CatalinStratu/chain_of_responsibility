package main

import "fmt"

// Functions
func add(a1, a2 int) int {
	res := a1 + a2
	fmt.Println("Result: ", res)
	return 0
}

// Main function
func main() {

	fmt.Println("Start")

	// Executes in LIFO order
	defer fmt.Println("End")

	for i := 0; i < 10; i++ {
		defer add(i, i)
	}
	//defer add(34, 56)
	//defer add(10, 10)
}
