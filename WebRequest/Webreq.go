package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	fmt.Println("Learning Web Request")
	res, err := http.Get("https://jsonplaceholder.typicode.com/todos/1")
	if err != nil {
		fmt.Println("There was an error while Getting the response")
		return
	}

	defer res.Body.Close()
	fmt.Printf("Type fo the response is %T", res)

	/* Read the response body */
	data, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println("Erro while reading the response body", err)
		return
	}

	fmt.Println("Response: ", string(data))
}
