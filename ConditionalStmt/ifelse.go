package main

import "fmt"

func main() {
	x := 10
	if x > 5 {
		fmt.Println("X is greater than 5")
	} else {
		fmt.Println("X is smaller than 5")
	}

	z := 10
	if z > 10 {
		fmt.Println("z is greater than 10")
	} else if z > 5 {
		fmt.Println("z is greater than 5")
	} else {
		fmt.Println("z is smaller than 5")
	}

	/* In GoLang the compiler automatically detects where the parenthesis are required and if not required then it removes the parenthesis automatically */

	y := 10
	if x < 5 && (y > 5 || z < 43) {
		fmt.Println("Hey, how are you?")
	} else {
		fmt.Println("Learn programming with hello world")
	}

	/* Here you can notice the difference between the above mentioned if else cases
	The earlier two case doesn't need any parenthesis, but the third case needs parenthesis as it's quite complex condition which is to be evaluated */
}
