package main

import (
	"fmt"
)

func main() {
	// cara mendeklarasikan slice dan langsung mengisinya
	fruits := []string{"apple", "banana", "orange", "watermelon"}

	// cara deklarasi ke dua slice kosong untuk mengisi dengan data dinamis menggunakan append
	// var fruits []string

	fruits = append(fruits, "coconut")
	fruits = append(fruits, "melon")
	fruits = append(fruits, "pinutes")

	for _, fruit := range fruits {
		fmt.Println("Fruit : ",fruit)
	}
}
