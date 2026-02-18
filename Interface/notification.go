package main

import "fmt"

type notifier interface {
	notify(msg string)
}

type notification struct{}

func (n notification) notify(str string) {
	fmt.Printf("Hello, you have a mail:%s\n", str)
}

type Sms struct{}

func (p Sms) notify(str string) {
	fmt.Printf("Hello, you have a sms:%s\n", str)
}

func main() {
	fmt.Println("Hello")
	nnewnoti := notification{}
	newSMS := Sms{}
	nnewnoti.notify("Hello")
	newSMS.notify("Hiii")
}
