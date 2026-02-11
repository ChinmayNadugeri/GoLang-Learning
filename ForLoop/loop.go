package main

import "fmt"

func main() {
	/* numbers := []int{25, 34, 58, 65, 45, 98} */

	/* for i := 0; i < len(numbers); i++ {
		fmt.Printf("Number at index %d is %d\n", i, numbers[i])
	} */

	/* Infinite loop */
	/* 	counter := 0
	 */ /* for {
		fmt.Println("Infinite loop")
		counter++

		if counter == 4 {
			break
		}
	} */

	/* For loop with range */

	/* for index, value := range numbers {
		fmt.Printf("Number at index %d is %d\n", index,value )
	} */

	/* For loop with range on string*/
	data := "Hello world"
	for index, value := range data {
		fmt.Printf("Number at index %d is %c\n", index, value)
	}
}
