package main

import (
	"fmt"
	"pertama/calculation"
)

func main() {
	fmt.Println("helloo world")
	result := calculation.Add(9, 9)
	// multi := calculation.Multiply(9, 9)

	fmt.Println(result)
	// fmt.Println(multi)
}
