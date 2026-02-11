package main

import (
	"fmt"
	"io/ioutil"
)

func main() {

	/* Step1 creating the file */
	/* file, err := os.Create("exmaple.txt")
	if err != nil {
		fmt.Println("There aws an error while creating the file")
		return
	}

	fmt.Println("File Created file successfully")

	defer file.Chdir() */

	/* Step2 Wrinting into  the file */
	/* content := "helo guyx, welcome to my YOutube Channel"
	byte, errors := io.WriteString(file, content+"\n")
	fmt.Println("Byte written id: ", byte)
	if errors != nil {
		fmt.Println("Error while writing the content into the file")
		return
	} */

	/* Processing the file */

	/* file,err:=os.Open("example.txt")
	if err != nil {
		fmt.Println("There was an error while opening the file")
		return
	}

	defer file.Close()
	*/
	/* Create a buffer to read the file content */
	/* buffer :=make([]byte,1024) */

	/* Read the content of the file into buffer*/

	/* for{
	n,err:=file.Read(buffer)
	if err==io.EOF{
		break
	}
	if err!=nil{
		fmt.Println("Error while reading the file",err)
		return
	} */
	/* Process the read content */
	/* 	fmt.Println(string(buffer[:n]))
	} */

	/* Read the content of the file using iotuil function */
	content, err := ioutil.ReadFile("example.txt")
	/* content, err := os.ReadFile("example.txt") */
	if err != nil {
		fmt.Println("there was an error while reading the file", err)
		return
	}

	fmt.Println(string(content))
}
