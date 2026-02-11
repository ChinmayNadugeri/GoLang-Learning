package main

import "fmt"

func main() {

	studentsgrade := make(map[string]int)

	studentsgrade["Chinmay"] = 85
	studentsgrade["Tanmay"] = 99
	studentsgrade["Prajwal"] = 95
	studentsgrade["Sourabh"] = 100
	studentsgrade["Koujalagi"] = 100

	fmt.Println("Marks of Chinmay: ", studentsgrade["Chinmay"])

	/* Updation of value */
	fmt.Println("Marks of Prajwal: ", studentsgrade["Prajwal"])
	studentsgrade["Prajwal"] = 100
	fmt.Println("Marks of Prajwal: ", studentsgrade["Prajwal"])

	/* Deletion of  value*/
	delete(studentsgrade, "Koujalagi")

	/* Checking existance */
	grades, exist := studentsgrade["Kpujalagi"]
	fmt.Println("Grades of Koujalagi: ", grades)
	fmt.Println("Koujalagi's existance: ", exist)

	/* Printing using for range */
	for index, value := range studentsgrade {
		fmt.Printf("Key is %s and grade is %d\n", index, value)
	}

	person := map[string]int{
		"ALice":   90,
		"Bob":     56,
		"Charlie": 45,
	}

	for index, value := range person {
		fmt.Printf("Key is %s and grade is %d\n", index, value)
	}

}
