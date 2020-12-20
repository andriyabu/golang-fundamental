package main

import (
	"fmt"
	"quiz/calculation"
)

func main() {
	fmt.Println("ini adalah fungsi perkalian")
	result := calculation.Multiply(9, 9)
	fmt.Println("9 x 9 adalah ", result)
}
