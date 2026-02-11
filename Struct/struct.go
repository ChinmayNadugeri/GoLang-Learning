package main

import "fmt"

type Person struct {
	FirstName string
	LastName  string
	Age       int
}

type Contact struct {
	Phone int
	Email string
}

type Address struct {
	HouseNo int
	Area    string
	State   string
}

/* Struct in Struct */

type Employee struct {
	Person_details Person
	Person_address Address
	Person_contact Contact
}

func main() {
	/* var stud1 person */
	/* fmt.Println("Chinmay person: ", stud1) */
	/* stud1.FirstName = "Chinmay"
	stud1.LastName = "Nadugeri"
	stud1.Age = 21 */
	/* fmt.Println("student 1: ", stud1) */

	/* stud2 := person{
		FirstName: "Prajwal",
		LastName:  "Karigoudar",
		Age:       21,
	} */
	/* fmt.Println("student 2 : ", stud2) */

	/* var stud3 = new(Person)
	stud3.FirstName = "Tanmay"
	stud3.LastName = "Tanmay"
	stud3.Age = 21 */

	/* fmt.Println("student 2 : ", stud3) */

	var employee1 Employee
	employee1.Person_details = Person{
		FirstName: "Chinmay",
		LastName:  "Nadugeri",
		Age:       21,
	}

	employee1.Person_address.Area = "Rajatgiri"
	employee1.Person_address.HouseNo = 47
	employee1.Person_address.State = "Karnataka"

	employee1.Person_contact.Phone = 8217415515
	employee1.Person_contact.Email = "chinmaynadugeri@gmail.com"

	fmt.Println(employee1)

}
