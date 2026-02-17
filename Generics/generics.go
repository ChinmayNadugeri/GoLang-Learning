package main

import "fmt"

func printSlice[T int | string | bool](items []T) {
	for _, item := range items {
		fmt.Println(item)
	}
}

/* func printSlice[T comparable](items []T) { //Here "Comparable keyword allows all datatypes in it."
	for _, item := range items {
		fmt.Println(item)
	}
} */

/* func printSlice[T comparable](items []T, V string) { //multiple generic types can be passed in it."
	for _, item := range items {
		fmt.Println(item)
	}
} */

// Using "any" keyword everytime is also not a good practice
/* func printSlice[T any](items []T) {
	for _, item := range items {
		fmt.Println(item)
	}
} */

//As we have used "any" keyword to access any type of the datatypes in the above function,
// its not a good practice for the coders. So we can also use an empty interface.
/* func printSlice[T interface{}](items []T) {
	for _, item := range items {
		fmt.Println(item)
	}
} */

/* func printStringSlice(items []string) {
	for _, item := range items {
		fmt.Println(item)
	}
} */

type stack[T any] struct { //Generics can also be used in struct like this.
	elements []T
}

func main() {

	myStack1 := stack[int]{
		elements: []int{1, 2, 3},
	}
	myStack2 := stack[string]{
		elements: []string{"Golang", "Typescript"},
	}
	fmt.Println(myStack1)
	fmt.Println(myStack2)
	/* nums := []int{1, 2, 3}
	names := []string{"Golang", "typescript"}
	vals := []bool{true, false, true}
	printSlice(nums)
	printSlice(names)
	printSlice(names)
	printSlice(vals) */

}
