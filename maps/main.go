package main

import "fmt"

func main() {
	m := make(map[string]int)
	m["Jose"] = 14
	m["Pepito"] = 20

	// Otra manera de declarar maps
	m2 := map[string]int{
		"Carlos": 13,
		"Saul":   34,
	}

	fmt.Println(m)
	fmt.Println(m2)

	// Recorrer un map
	for i, v := range m {
		fmt.Println(i, v)
	}

	// Encontrar un valor
	value := m["Jose"]
	fmt.Println(value)

	// Valor no encontrado
	value1, ok := m["NotFound"]
	fmt.Println(value1, ok)
}
