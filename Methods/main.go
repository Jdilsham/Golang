package main

import "fmt"

func main() {
	fmt.Println("Welcome to Methods")

	JD := person{
		Name:   "Janitha",
		Age:    25,
		Status: true,
	}
	fmt.Println("Name:", JD.Name)
	fmt.Println("Age:", JD.Age)
	JD.greet()
	JD.NewMail()

	fmt.Println("Email:", JD.Email)
	//Email functions provide a copy, it not changes actual value
	//That's we use pointers
}

type person struct {
	Name   string
	Age    int
	Status bool
	Email  string
}

func (p person) greet() {
	fmt.Println("Is you active:", p.Status)
}

func (p person) NewMail() {
	p.Email = "jd@gmail.com"
	fmt.Println("Email:", p.Email)
}
