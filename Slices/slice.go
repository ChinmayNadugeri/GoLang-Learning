package main

import "fmt"

func main() {

	numbers := []int{1, 3, 3, 4, 5, 6}
	/* We can also append the numbers dynamically to the slice */
	numbers = append(numbers, 7, 8, 9, 10)
	fmt.Println("Numbers: ", numbers)
	fmt.Printf("Numbers has data type of: %T\n", numbers)
	fmt.Println("Lnegth of numbers is: ", len(numbers))

	/* We can also initialize the slice using make function  */
	nos := make([]int, 3, 5) /* In this function the initial parameter is the slice to be created,
	   The second parameter is the length(No of elements in ths slice) of the slice
	   The third parameter is the capacity(capacity of the slice{i.e, how many elements can be stored in the slice}) of the slice*/

	/* The capacity of the slice doubles once the length of the slice exceeds capacity of the slice*/
	fmt.Println(nos)
	fmt.Println("Length of nos:", len(nos))
	fmt.Println("Length of nos:", cap(nos))
}
