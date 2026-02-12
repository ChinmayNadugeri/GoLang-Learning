package main

import (
	"fmt"
)

// The Variadic function in go is almost similar to the spread operator in javaScript
func sum(nums ...int) int {
	total := 0
	for _, num := range nums {
		total = total + num
	}
	return total
}

func main() {
	res := sum(1, 2, 3, 4, 5)
	//we can pass the slice as the parameter to this sum function
	/*nums:=[]int{1,2,3,4,5}
	res:=sum(nums...)*/
	fmt.Println("Result is: ", res)
}
