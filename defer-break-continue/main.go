package main

import "fmt"

func main() {
	// Defer: ejecuta código al final de que muera la ejecución
	defer fmt.Println("Hola")
	fmt.Println("Mundo")

	// Continue y break
	// El continue es para omitir una iteración de un bucle, dada una condición
	for i := 0; i < 10; i++ {
		fmt.Println(i)

		// Continue
		if i == 2 {
			fmt.Println("Es 2")
			continue
		}

		// Break
		if i == 8 {
			fmt.Println("Break")
			break
		}
	}
}
