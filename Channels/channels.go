package main

import (
	"fmt"
	/* "math/rand"
	"time" */)

// Sending
/* func processNum(numChan chan int) {

	for num := range numChan {
		fmt.Println("Processing number", num)
		time.Sleep(time.Second)
	}

} */

// recieve
func sum(result chan int, num1 int, num2 int) {
	numResult := num1 + num2
	result <- numResult
}

func main() {

	result := make(chan int)

	go sum(result, 4, 5)
	res := <-result //blocking
	fmt.Println(res)
	/* numChan := make(chan int)

	go processNum(numChan)

	for {
		numChan <- rand.Intn(100)
	}
	*/
	/* messageChan := make(chan string)

	messageChan <- "Ping" //blocking

	msg:= <- messageChan //Receiving
	fmt.Println(msg) */
}
