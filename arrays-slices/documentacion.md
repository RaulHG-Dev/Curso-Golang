# Arrays y Slices

> [!NOTE]  
> La diferencia principal entre los arrays es que estos tienen una longitud fija e invariable y deben declarase especificándola, mientras que los Slices tienen una longitud variable y no hay que especificarla en la declaración.

* Ejemplo:
```go
package main

import "fmt"

func main() {
	/**
	NOTA: La diferencia principal entre los arrays es que estos tienen una longitud fija e invariable y deben declarase especificandola,
	mientras que los Slices tienen una longitud variable y no hay que especificarla en la declaración
	*/
	// Array: En go los array son inmutables
	var array [4]int

	array[0] = 1
	array[1] = 2
	fmt.Println(array, len(array), cap(array))

	// Slice
	slice := []int{0, 1, 2, 3, 4, 5, 6}
	fmt.Println(slice, len(slice), cap(slice))

	// Métodos en el slice
	fmt.Println(slice[0])
	fmt.Println(slice[:3])
	fmt.Println(slice[2:4])
	fmt.Println(slice[4:])

	// Append
	slice = append(slice, 7)
	fmt.Println(slice)

	//  Append nueva lista
	newSlice := []int{8, 9, 10}
	slice = append(slice, newSlice...)
	fmt.Println(slice)
}
```
