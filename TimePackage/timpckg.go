package main

import (
	"fmt"
	"time"
)

func main() {
	currenttime := time.Now()
	/* fmt.Println("Current time: ", currenttime)*/
	formatted := currenttime.Format("02/01/2006, Monday")
	fmt.Println("Fomatted date and time: ", formatted)

	layout_str := "02-01-2006" /* This si the standard format whih should be provided for formatting the time*/
	datatsr := "25-11-2030"
	foramtted_time, _ := time.Parse(layout_str, datatsr)
	fmt.Println("Foramted time: ", foramtted_time)

	/*add 1 more day in your current day*/
	new_date := currenttime.Add(24 * time.Hour)
	fmt.Println("New_date: ", new_date)
	formatted_newdate := new_date.Format("02/01/2006")
	fmt.Println("formatted_newdate time: ", formatted_newdate)
}
