package main

import (
	"encoding/json"
	"fmt"
)

type Person struct {
	Name    string `json:"name"`
	Age     int    `json:"age"`
	IsAdult bool   `json:"is_adult"`
}

func main() {

	fmt.Println("We are learning JSON in Go")
	person := Person{"John", 34, true}
	/* fmt.Println("Person Data is : ", person) */

	/* Converting person into json Encoding {Marshaling} */
	jsonData, err := json.Marshal(person)
	if err != nil {
		fmt.Println("Encountered error while marshalling")
		return
	}
	fmt.Println("Person data is : ", string(jsonData))

	/* Decoding {unmarshalling} */
	var personData Person
	err = json.Unmarshal(jsonData, &personData)
	if err != nil {
		fmt.Println("Error while Unmarshalling", err)
		return
	}
	fmt.Println("Person data after unmarshalling is:", personData)
}
