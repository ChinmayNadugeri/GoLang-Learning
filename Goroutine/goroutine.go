package main

import (
	"fmt"
	"time"
)

func SayHello() {
	fmt.Println("Hello world!")
	time.Sleep(2000 * time.Millisecond) //Simulating some work
	fmt.Println("SayHello function executed successfully ")

}

func SayHi() {
	fmt.Println("Hi Chinmay")
	time.Sleep(1000 * time.Millisecond) //Simulating some work
	fmt.Println("SayHi function executed successfully ")
}
func main() {
	fmt.Println("Learning goroutine")
	go SayHello()
	go SayHi()

	//Wait for a moment to allow the goroutine to finish
	time.Sleep(800 * time.Millisecond)
}
