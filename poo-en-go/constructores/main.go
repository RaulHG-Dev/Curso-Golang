package main

import "fmt"

type Employee struct {
	id       int
	name     string
	vacation bool
}

func NewEmployee(id int, name string, vacation bool) *Employee {
	return &Employee{
		id:       id,
		name:     name,
		vacation: vacation,
	}
}

func main() {
	// Manera 1
	e := Employee{}
	fmt.Printf("%v", e)
	// Manera 2
	e2 := Employee{
		id:       1,
		name:     "Carlos",
		vacation: true,
	}
	fmt.Println(e2)
	// Manera 3
	e3 := new(Employee)
	fmt.Println(*e3)
	e3.id = 2
	e3.name = "Roberto"
	fmt.Println(*e3)
	// Manera 4
	e4 := NewEmployee(1, "Ruben", false)
	fmt.Println(*e4)

}
