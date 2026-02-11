package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	//Using Scanf//
	// var name string
	// fmt.Println("Hey, what is your name?")
	// fmt.Scan(&name)
	// fmt.Println("Hello Mr.", name)

	//Using buffer reader//
	reader := bufio.NewReader(os.Stdin)
	name, _ := reader.ReadString('\n')
	fmt.Println("Hello Mr.", name)

}
