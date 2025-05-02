package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func main() {
	fmt.Println("Post Request")
	PerformpostRequest()
}

func PerformpostRequest() {

	const myurl = "http://localhost:8080/hello"

	requestBody := strings.NewReader(`{
		"name": "John",
		"address": "123 Main St" 
	}`)

	response, err := http.Post(myurl, "application/json", requestBody)

	if err != nil {
		panic(err)
	}

	defer response.Body.Close()

	fmt.Println(response.Status)
	fmt.Println(response.ContentLength)

	content, _ := ioutil.ReadAll(response.Body)
	fmt.Println(string(content))

}
