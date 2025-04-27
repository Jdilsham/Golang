package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main() {
	fmt.Println("Welcome to Files")

	content := "Janitha Dilsham Wanni Arachchi"

	file, err := os.Create("file.txt")

	//if err != nil {
	//	panic(err)
	//}

	length, err := io.WriteString(file, content)

	checkNilErr(err)

	fmt.Println("Length of file is ", length)
	defer file.Close()

	readFile("file.txt")
}

func readFile(filename string) {
	fileinByte, err := ioutil.ReadFile(filename)
	checkNilErr(err)

	fmt.Println(string(fileinByte)) //File is displayed as bytecode, using string we can get text
}

func checkNilErr(err error) {
	if err != nil {
		panic(err)
	}
}
