package main

import "fmt"

func add(a, b int) int {
	return a + b
}
func main() {

	/* Whenever the defer fucntion is used in a program its role is to do that work in the last of the main function. i.e, as mentioned in the below program, the line with the defer keyword will be executed in the last of the main function */
	fmt.Println("Starting of the programming")
	data := add(3, 4)
	defer fmt.Println("Data is : ", data)
	fmt.Println("Middle of the programming")
	fmt.Println("End of the programming")

	/* Whenever more than one defer keyword is used in a program the defer works follows Stack execution. i.e, the last statement with the defer will be executed first [LIFO]  */
}
