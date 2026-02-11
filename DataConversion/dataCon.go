package main

import (
	"fmt"
	"strconv"
)

func main() {
	num := 10
	fmt.Println("Value of num is: ", num)
	fmt.Printf("DataType of num is %T \n", num)

	var data float64 = float64(num)
	fmt.Println("Value of data is: ", data)
	fmt.Printf("DataType of data is %T \n", data)

	num = 123
	str := strconv.Itoa(num) /* string conversion "Itoa => I refers to Integer", "a => a refers to string" */
	fmt.Println("Value of str is: ", str)
	fmt.Printf("DataType of str is %T \n", str)

	num_string := "12345"
	num_int, _ := strconv.Atoi(num_string) /* string conversion "Atoi => A refers to string", "i => i refers to Integer" */
	fmt.Println("Value of num1 is: ", num_int)
	fmt.Printf("DataType of num1 is %T \n", num_int)

	num_string1 := "3.14"
	num_float, _ := strconv.ParseFloat(num_string1, 64) /* string conversion  str=>float64" */
	fmt.Println("Value of num_float is: ", num_float)
	fmt.Printf("DataType of num_float is %T \n", num_float)
}
