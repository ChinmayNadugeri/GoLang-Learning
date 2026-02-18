package main

import "fmt"

type paymenter interface {
	pay(amount float32)
	refund(amount float32, account string)
}

type payment struct {
	gateway paymenter
}

/* Open Close Principle: The class and methods should be open for extension and close for modification*/
func (p payment) makePayment(amount float32) {
	/* razorpayPaymentGW := razorpay{} */
	/* razorpayPaymentGW.pay(amount) */
	p.gateway.pay(amount)
}

func (p payment) refund(amount float32, account string) {
	fmt.Println(amount, "refund successfully done for acc:", account)
}

type razorpay struct{}

func (r razorpay) pay(amount float32) {
	fmt.Println("Making payment using razorpay", amount)
}

type stripe struct{}

func (s stripe) pay(amount float32) {
	fmt.Println("Making payment using stripe", amount)
}

type fakePayment struct{}

func (f fakePayment) pay(amount float32) {
	fmt.Println("making payment using fakepayment gateway for  testng purpose")
}

type paypal struct{}

func (p paypal) pay(amount float32) {
	fmt.Println("Making payment using Paypal", amount)
}

func (p paypal) refund(amount float32, account string) {
	fmt.Println(amount, "refund successfully done for acc:", account)
}

func main() {
	/* stripePaymentGW := stripe{} */
	/* razorpayPaymentGW:=razorpay{} */
	/* fakeGW:=fakePayment{} */
	paypalGW := paypal{}
	newPayment := payment{
		gateway: paypalGW,
	}
	newPayment.makePayment(2000)
	newPayment.refund(1000, "74180100012735")
}
