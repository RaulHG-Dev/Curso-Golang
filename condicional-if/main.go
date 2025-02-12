package main

import (
	"fmt"
	"log"
	"strconv"
)

func main() {
	valor1 := 1
	valor2 := 2

	if valor1 == valor2 {
		fmt.Println("Es 1")
	} else {
		fmt.Println("No es 1")
	}

	// Con AND
	if valor1 == 1 && valor2 == 2 {
		fmt.Println("Es verdad")
	}

	// Con OR
	if valor1 == 0 || valor2 == 2 {
		fmt.Println("Es verdad, OR")
	}

	// Convertir texto a número
	value, err := strconv.Atoi("53")

	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Println("Value:", value)
	}
}
