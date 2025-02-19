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
	// endDate string
}

func (ftEmployee FullTimeEmployee) getMessage() string {
	return "Full Time Employee"
}

type TemporaryEmployee struct {
	Person
	Employee
	// taxRate int
}

func (tEmployee TemporaryEmployee) getMessage() string {
	return "Temporary Employee"
}

type PrintInfo interface {
	getMessage() string
}

func getMessage(p PrintInfo) {
	fmt.Println(p.getMessage())
}

func main() {
	ftEmployee := FullTimeEmployee{}

	ftEmployee.name = "Carlos"
	ftEmployee.age = 23
	ftEmployee.id = 5

	fmt.Println(ftEmployee)

	tEmployee := TemporaryEmployee{}
	getMessage(tEmployee)
	getMessage(ftEmployee)
}
