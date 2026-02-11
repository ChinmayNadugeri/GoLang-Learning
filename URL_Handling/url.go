package main

import (
	"fmt"
	"net/url"
)

func main() {
	/* fmt.Println("Learning URL Handling") */
	myurl := "https://lamacodetech.in/contact/"

	/* Converting the normal string url to actual URL format */
	ParsedURL, err := url.Parse(myurl)
	if err != nil {
		fmt.Println("Error encountered while parsing the url")
		return
	}

	/* fmt.Printf("Type of ParsedURL is %T", ParsedURL) */
	fmt.Println("Scheme of ParsedURL is ", ParsedURL.Scheme)
	fmt.Println("Host of ParsedURL is ", ParsedURL.Host)
	fmt.Println("Path of ParsedURL is ", ParsedURL.Path)
	fmt.Println("RawQuery of ParsedURL is ", ParsedURL.RawQuery)

	/* Modifying the parsed URL */
	ParsedURL.Path = "/newPath"
	ParsedURL.RawQuery = "username=Chinmay"

	newUrl := ParsedURL.String()

	fmt.Println("new URL: ", newUrl)

}
