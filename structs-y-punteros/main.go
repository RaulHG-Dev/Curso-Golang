package main

import "fmt"

type pc struct {
	ram   int
	disk  int
	brand string
}

func (myPC pc) ping() {
	fmt.Println(myPC.brand, "Pong")
}

// The `duplicateRAM` method defined for the `pc` struct is a method that operates on a pointer
// receiver (`*pc`). This method takes the current `pc` instance, multiplies its RAM value by 2, and
// updates the RAM value of the instance itself. Since the receiver is a pointer to the `pc` struct,
// any changes made to the struct within this method will be reflected in the original struct outside
// the method.
func (myPC *pc) duplicateRAM() {
	myPC.ram = myPC.ram * 2
}

func main() {
	a := 50
	// Puntero de a
	b := &a

	fmt.Println(a, b)
	// Acceder al valor de almacenado en la dirección del puntero
	fmt.Println(*b)

	myPc := pc{ram: 16, disk: 200, brand: "msi"}
	fmt.Println(myPc)

	myPc.ping()

	fmt.Println(myPc)

	myPc.duplicateRAM()
	fmt.Println(myPc)
}
