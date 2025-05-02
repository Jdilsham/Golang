package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func main() {
	fmt.Println("Get Request in golang")

	performGetRequest()
}

func performGetRequest() {
	const myurl = "http://localhost:8080/hello"

	response, err := http.Get(myurl)
	if err != nil {
		panic(err)
	}

	defer response.Body.Close()

	fmt.Println("Response Status: ", response.Status)
	fmt.Println("Content Length: ", response.ContentLength)

	var responseString strings.Builder
	content, err := ioutil.ReadAll(response.Body)
	byteCount, _ := responseString.Write(content)
	if err != nil {
		panic(err)
	}

	fmt.Println("Byte Count: ", byteCount)
	fmt.Println("Response Body: ", responseString.String())
	//fmt.Println("Response Body: ", string(content))1.
	//    - `https://www.google.com`

}
