package main

import (
	"fmt"
	"net/url"
)

const myurl string = "https://lco.dev:3000/learn?coursename=reactjs&paymentid=123abc"

func main() {
	fmt.Println("Welcome to handling URL in golang")
	fmt.Println(myurl)

	result, err := url.Parse(myurl)

	if err != nil {
		panic(err)
	}

	fmt.Println(result)

	fmt.Println(result.Host)
	fmt.Println(result.Scheme)
	fmt.Println(result.Path)
	fmt.Println(result.Port())
	fmt.Println(result.RawQuery)
}
