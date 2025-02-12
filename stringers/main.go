package main

import "fmt"

type pc struct {
	ram   int
	brand string
	disk  int
}

// This function is a method defined on the `pc` struct. It overrides the default `String()` method for
// the `pc` struct type. When you call `fmt.Println(myPc)`, it will use this custom `String()` method
// to format the output string for the `pc` struct instance `myPc`.
func (myPc pc) String() string {
	return fmt.Sprintf("Tengo %d GB de RAM, %d GB de Disco y es una %s", myPc.ram, myPc.disk, myPc.brand)
}

func main() {
	myPc := pc{ram: 16, brand: "msi", disk: 100}

	fmt.Println(myPc)
}
