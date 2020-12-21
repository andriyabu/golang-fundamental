package main

import (
	"fmt"
)

func main() {
	// deklarasi variable langsung
	// numberA := 5
	// numberB := &numberA // numberB diisi dengan alamat memory numberA /referensiasi

	var numberA int = 5
	var numberB *int = &numberA // numberB diisi dengan alamat memory numberA /referensiasi

	fmt.Println(numberA)  // 5
	fmt.Println(numberB)  // 0xc00010c000
	fmt.Println(*numberB) // 5

	*numberB = 20

	fmt.Println(numberA)  // 10
	fmt.Println(numberB)  // 0xc00010c000
	fmt.Println(*numberB) // 10

}
