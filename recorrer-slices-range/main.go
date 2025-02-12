package main

import (
	"fmt"
	"strings"
)

func isPalindromo(text string) {
	var textReverse string
	textToLowerCase := strings.ToLower(text)
	for i := len(text) - 1; i >= 0; i-- {
		textReverse += string(textToLowerCase[i])
	}

	if textReverse == textToLowerCase {
		fmt.Println("Es palindromo")
	} else {
		fmt.Println("No es un palindromo")
	}
}

func main() {
	slice := []string{"hola", "que", "hace"}

	for i, valor := range slice {
		fmt.Println(i, valor)
	}

	isPalindromo("ama")
	isPalindromo("Ama")
	isPalindromo("casas")
}
