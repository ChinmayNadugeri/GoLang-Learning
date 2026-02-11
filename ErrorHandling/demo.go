package main

import "fmt"

func divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, fmt.Errorf("denominator cannot be zero")
	}
	return a / b, nil
}

/* We can also use string for returning the error
func divide(a, b float64) (float64, string) {
	if b == 0 {
		return 0, "denominator cannot be zero"
	}
	return a / b, "nil"
}
*/

func main() {

	var a float64
	var b float64
	var ans float64
	fmt.Println("Enter a: ")
	fmt.Scan(&a)
	fmt.Println("Enter b: ")
	fmt.Scan(&b)
	ans, _ = divide(a, b)
	/* Herer "_" is used as a return parameter for error, whic is ignored by the compiler  */

	/* Another method to define err instead of using "_" */
	/* ans, err:=divide(a,b)
	if(err!=nil){
		fmt.Println("Error handling")
	} */
	fmt.Println("Answer is : ", ans)

}
