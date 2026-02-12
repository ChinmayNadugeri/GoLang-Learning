package main

import (
	"fmt"
	"sync"
)

func worker(i int, wg *sync.WaitGroup) {
	defer wg.Done() //Sends a signal to the WaitGroup that it has done the all work
	fmt.Printf("Worker %d started\n", i)
	//Some work is happening
	fmt.Printf("Worker %d ended\n", i)
}
func main() {
	var wg sync.WaitGroup
	fmt.Println("Learning Sync WaitGroup in GoLang")
	//starts 3 workers routine
	for i := 1; i <= 3; i++ {
		wg.Add(1) //Increment the wait group counter
		go worker(i, &wg)
	}

	wg.Wait()
	fmt.Println("Task successfully executed")
}
