package main

import "fmt"

func modifyValueByRef(num *int) {
	*num = *num * 20
}

func main() {
	num := 2

	ptr := &num

	fmt.Println("Num has value: ", num)
	fmt.Println("Pointer contains: ", ptr)

	/* Retrieving value using pointer */
	fmt.Println("Data contains pointer through: ", *ptr)

	/* If a pointer ins only declared and not assigned any value it stores "nil" as initial value */
	var ptr1 *int
	if ptr1 == nil {
		fmt.Println("Pointer is not assigned")
	}

	value := 5
	modifyValueByRef(&value)
	fmt.Println("Value contains: ", value)

}
