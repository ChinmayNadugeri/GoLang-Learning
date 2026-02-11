package main

import "fmt"

func main() {
	fmt.Println("We are learning Array")

	var name [5]string
	name[0] = "Chinmay"
	name[1] = "Prajwal"
	name[2] = "Tanmay"

	fmt.Println("Name of the person is: ", name)
}

/* In GoLang whenever an array is initialized, the values are assigned as
as "0" for integer arrays, "false" for boolean arrays, "" for string arrays
When an empty array of string is printed it doesn't show any of the output as it contains nothing in it.
To print the empty string array and confirm that its is empty we use "%q* formatter to show quoted space in an empty stir==ring array

fmt.Printf("The String array is %q\n: ",name)

*/
