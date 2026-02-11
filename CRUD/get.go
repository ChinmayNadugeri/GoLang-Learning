package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

type Todo struct {
	UserID    int    `json:"userid"`
	ID        int    `json:"id"`
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
}

func performGetRequest() {

	res, err := http.Get("https://jsonplaceholder.typicode.com/todos/1")
	if err != nil {
		fmt.Println("Error while getting the data", err)
		return
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		fmt.Println("Error encountered while getting the data", err)
		return
	}

	/* data, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println("Error encountered while reading", err)
		return
	}
	fmt.Println("The data is: ", string(data)) */

	var todo Todo
	err = json.NewDecoder(res.Body).Decode(&todo)
	if err != nil {
		fmt.Println("Error whil decoding", err)
		return
	}

	fmt.Println("Todo: ", todo)
	fmt.Println("GET Response: ", res.Status)
	/* The second method directly stores the json value of todo into the struct todo*/
}

func performPostRequest() {
	todo := Todo{
		UserID:    23,
		Title:     "Chinmay",
		Completed: true,
	}

	//Convert  the todo to json
	jsonData, err := json.Marshal(todo)
	if err != nil {
		fmt.Println("Error while marshalling", err)
		return
	}

	//convert json to string for passing it as a parameter in "Post()" function
	jsonString := string(jsonData)

	//convert the string to reader
	jsonreader := strings.NewReader(jsonString)

	myurl := "https://jsonplaceholder.typicode.com/todos"
	res, err := http.Post(myurl, "application/json", jsonreader)
	if err != nil {
		fmt.Println("Error while sending request", err)
		return
	}

	defer res.Body.Close()

	/* data, _ := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println("Error while reading the reuskt body", err)
		return
	}
	fmt.Println("Response: ", string(data)) */

	fmt.Println("POST Response Status: ", res.Status)

}

func performUpdateRequest() {
	todo := Todo{
		UserID:    2351012,
		Title:     "Chinmay Nadugeri",
		Completed: false,
	}

	//Convert  the todo to json
	jsonData, err := json.Marshal(todo)
	if err != nil {
		fmt.Println("Error while marshalling", err)
		return
	}

	//convert json to string for passing it as a parameter in "Post()" function
	jsonString := string(jsonData)

	//convert the string to reader
	jsonreader := strings.NewReader(jsonString)

	const myurl = "https://jsonplaceholder.typicode.com/todos/1"

	//create PUT Request
	req, err := http.NewRequest(http.MethodPut, myurl, jsonreader)
	if err != nil {
		fmt.Println("Error while creating PUT request", err)
		return
	}

	req.Header.Set("Content-type", "application/json")

	//send the request
	client := http.Client{}
	res, err := client.Do(req)
	if err != nil {
		fmt.Println("Error while sending requesting", err)
		return
	}

	defer res.Body.Close()

	data, _ := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println("Error while reading the reuskt body", err)
		return
	}
	fmt.Println("Response: ", string(data))
	fmt.Println("PUT request Response: ", res.Status)
}

func performDeleteRequest() {
	const myurl = "https://jsonplaceholder.typicode.com/todos/1"

	//create DELETE request
	req, err := http.NewRequest(http.MethodDelete, myurl, nil)
	if err != nil {
		fmt.Println("Error while creating DELETE request", err)
		return
	}

	//send the request
	client := http.Client{}
	res, err := client.Do(req)
	if err != nil {
		fmt.Println("Error while sending requesting", err)
		return
	}

	defer res.Body.Close()
	fmt.Println("DELETE Response : ", res.Status)

}
func main() {
	fmt.Println("Learning CRUD in Go")
	performGetRequest()
	performPostRequest()
	performUpdateRequest()
	performDeleteRequest()
}
