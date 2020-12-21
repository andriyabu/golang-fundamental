package main

import (
	"fmt"
)

func main() {
	numberA := 5
	numberB := &numberA // numberB diisi dengan alamat memory numberA

	fmt.Println(numberA)  // 5
	fmt.Println(numberB)  // 0xc00010c000
	fmt.Println(*numberB) // 5

	*numberB = 10

	fmt.Println(numberA)  // 10
	fmt.Println(numberB)  // 0xc00010c000
	fmt.Println(*numberB) // 10

}
