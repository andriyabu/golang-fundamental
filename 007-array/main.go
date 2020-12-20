package main

import (
	"fmt"
)

func main() {
	// deklarasi variabel array model pertama
	var fruits [5]string

	// mengisi array
	fruits[0] = "apple"
	fruits[1] = "grape"
	fruits[2] = "banana"
	fruits[3] = "coconut"
	fruits[4] = "lemon"

	fmt.Println(fruits)
	length := len(fruits)
	fmt.Println("panjang array :",length)
	fmt.Println()

	var language = [...]string{"Ruby", "Golang", "Java", "C++", "Python"}

	for index, value := range language {
		fmt.Println("index :", index, " language: ", value)
	}
}
