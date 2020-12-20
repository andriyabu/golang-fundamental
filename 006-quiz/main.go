package main

import (
	"fmt"
)

func main() {

	languange := "Golang is the best"

	for _, value := range languange {

		// menampilkan index genap
		// if index%2 == 0 {
		// 	fmt.Println(string(value))
		// }

		letterString := string(value)

		switch letterString {
		case "a", "i", "u", "e", "o":
			fmt.Println("Letter :", letterString)
		}
	}
}
