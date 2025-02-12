package main

import (
	pk "CursoGo/modificadores-acceso-structs/mypackage"
	"fmt"
)

func main() {
	var myCar pk.CarPublic

	myCar.Brand = "Ferrari"
	myCar.Year = 2021
	fmt.Println(myCar)

	pk.PrintMessage("Hola a Todos!")
}
