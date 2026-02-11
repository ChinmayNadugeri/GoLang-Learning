package main

import "fmt"

func add(x, y int) int {
	return x + y
}

func multiply(x, y int) int {
	return x * y
}

func main() {

	var a int
	var b int
	fmt.Println("Enter number a: ")
	fmt.Scan(&a)
	fmt.Println("Enter number b: ")
	fmt.Scan(&b)
	fmt.Printf("Sum of %d and %d is: %d \n", a, b, add(a, b))
	fmt.Printf("Product of %d and %d is: %d \n", a, b, multiply(a, b))
}
