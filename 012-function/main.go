package main

import (
	"fmt"
)

func main()  {
	result := hello("Saya sedang")
	fmt.Println(result)
}

func hello(sentence string) string  {
	mySentence := sentence + " Golang Fundamental"
	return mySentence
}