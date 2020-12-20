package main

import (
	"fmt"
)

func main() {
	//deklarasi variable myMap bertipe map dengan key string dan value string
	var myMap map[string]string
	//inisialisasi variable myMap bisa kosong
	// myMap = map[string]string{}
	// atau langsung diberikan nilai awal
	myMap = map[string]string{"nama": "andri yabu", "alamat": "perum mega regency"}

	// menambah data baru dalam map
	myMap["phone"] = "085623827832"

	fmt.Println(myMap)

	for index, value := range myMap {
		fmt.Println(index,": ",value)
	}
}
