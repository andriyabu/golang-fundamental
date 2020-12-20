package main

import (
	"fmt"
)

func main() {

	// for looping
	for i := 0; i <= 10; i++ {
		fmt.Println("belajar golang", i)
	}

	fmt.Println("=======================================")

	// for looping like while looping

	j := 1
	for j <= 10 {
		fmt.Println("Golang is the best", j)
		j++
	}

	fmt.Println("=======================================")

	languange := "Golang is the best programming language"

	for _, letter := range languange {
		fmt.Println("Letter", string(letter))
	}
}
