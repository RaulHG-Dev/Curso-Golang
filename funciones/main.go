package main

import "fmt"

func normalFunction(message string) {
	fmt.Println(message)
}

func tripleArguments(a, b int, c string) {
	fmt.Println(a, b, c)
}

func returnValue(a int) int {
	return a * 2
}

func dobleReturn(a int) (c, b int) {
	return a, a * 2
}

func main() {
	normalFunction("Hola mundo")
	tripleArguments(1, 2, "hola")

	value := returnValue(2)
	fmt.Println("Value:", value)

	value1, value2 := dobleReturn(2)
	fmt.Println("value1, value2:", value1, value2)

	// Ignorar alguno de los valores devueltos
	_, value2 = dobleReturn(3)
	fmt.Println("value2:", value2)

}
