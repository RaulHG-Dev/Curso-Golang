package main

import "fmt"

type Person struct {
	name string
	age  int
}

type Employee struct {
	id int
}

type FullTimeEmployee struct {
	Person
	Employee
}

func main() {
	ftEmployee := FullTimeEmployee{}

	ftEmployee.name = "Carlos"
	ftEmployee.age = 23
	ftEmployee.id = 5

	fmt.Println(ftEmployee)
}
