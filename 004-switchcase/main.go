package main

import (
	"fmt"
)

func main() {
	number := 2

	switch number {
	case 5:
		fmt.Println("Lima")
	case 6:
		fmt.Println("Enam")
	case 7:
		fmt.Println("Tujuh")
	case 8:
		fmt.Println("Delapan")
	case 9:
		fmt.Println("Sembilan")
	default:
		fmt.Println("DEFAULT")

	}
}
