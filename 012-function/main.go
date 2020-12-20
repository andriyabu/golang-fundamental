package main

import (
	"fmt"
)

func main() {
	result := hello("Saya sedang")
	fmt.Println(result)

	addresult := add(5, 5)
	fmt.Println(addresult)

	luas, keliling := calculate(8, 2)

	fmt.Println(luas)
	fmt.Println(keliling)

	kali, bagi := kalibagi(4, 2)

	fmt.Println(kali)
	fmt.Println(bagi)
}

func hello(sentence string) string {
	mySentence := sentence + " Golang Fundamental"
	return mySentence
}

func add(number, numberTwo int) int {
	return number + numberTwo
}

// multiple return
func calculate(panjang, lebar int) (int, int) {
	luas := panjang * lebar
	keliling := 2 * (panjang + lebar)

	return luas, keliling
}

//predefine function
func kalibagi(number, numberTwo int) (kali int, bagi int) {
	kali = number * numberTwo
	bagi = number / numberTwo

	return
}
