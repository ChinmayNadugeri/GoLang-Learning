package main

import (
	"fmt"
	"strings"
)

func main() {
	/* Split function in strings */
	data := "apple,orange,banana"
	parts := strings.Split(data, ",")
	fmt.Println(parts)

	/* Count function in strings */
	str := "one two three four two two six"
	count := strings.Count(str, "two")
	fmt.Println("Count is: ", count)

	/* trim function in strings */
	str1 := "       Hellloooo guyzz!     "
	trimmed := strings.Trim(str1, " ")
	fmt.Println(trimmed)
	/* This  trim function trims off the extra spaces from the starting of the string and at the end  */

	/* String concatination function */
	Firstname := "Chinmay"
	Lastname := "Nadugeri"
	FullName := strings.Join([]string{Firstname, Lastname}, " ")
	fmt.Println("Full name is: ", FullName)
}
