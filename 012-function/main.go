package main

import (
	"fmt"
)

func main()  {
	result := hello("Saya sedang")
	fmt.Println(result)

	addresult := add(5,5)
	fmt.Println(addresult)
}

func hello(sentence string) string  {
	mySentence := sentence + " Golang Fundamental"
	return mySentence
}

func add(number, numberTwo int) int  {
	return number + numberTwo
}