package mypackage

import "fmt"

// CarPublid Car con acceso público
type CarPublic struct {
	Brand string
	Year  int
}

type carPrivate struct {
	brand string
	year  int
}

func PrintMessage(text string) {
	fmt.Println(text)
}
