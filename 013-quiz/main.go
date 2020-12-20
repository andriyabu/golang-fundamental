package main

import (
	"errors"
	"fmt"
)

func main() {
	scores := []int{10, 5, 8, 9, 7}

	total := sum(scores)

	fmt.Println("Total : ", total)

	hasil, err := calculate(12, 6, ";")

	fmt.Println(hasil)
	if err != nil {
		fmt.Println(err)	
	}
	

}

func sum(arr []int) (sum int) {
	for _, value := range arr {
		sum += value
	}

	return
}

func calculate(number int, numberTwo int, operator string) (result int, err error) {
	switch operator {
	case "+":
		result = number + numberTwo
	case "-":
		result = number - numberTwo
	case "*":
		result = number * numberTwo
	case "/":
		result = number / numberTwo
	default:
		err = errors.New("Operator tidak dikenal")
	}

	return
}
