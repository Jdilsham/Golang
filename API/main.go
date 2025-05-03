package main

import "fmt"

// Model for course - file
type Course struct {
	CourseId    string  `json:"course_id"`
	CourseName  string  `json:"course_name"`
	CoursePrice float64 `json:"course_price"`
	Author      *Author `json:"author"`
}

type Author struct {
	FullName string `json:"full_name"`
	Website  string `json:"website"`
}

// fake DB
var courses []Course

//middleware, helper - file

func IsEmpty(c *Course) bool {
	return c.CourseId == "" && c.CourseName == ""
}

func main() {
	fmt.Println("Welcome to API")

}
