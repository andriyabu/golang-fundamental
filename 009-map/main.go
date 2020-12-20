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

	for key, value := range myMap {
		fmt.Println(key, ": ", value)
	}

	fmt.Println("======================")
	// menghapus map dengan fungsi delete(namaMap,key)
	delete(myMap, "alamat")

	for key, value := range myMap {
		fmt.Println(key, ": ", value)
	}

	fmt.Println("====================================")
	//mengecek apakah ada atau tidak key tertentu dalam map
	//  value, boolean = namaMap[key]
	value, isAvailable := myMap["alamat"] // false karena alamat sudah di hapus

	fmt.Println("is available :", isAvailable) // true kalau ada dan false kalau kosong
	fmt.Println(value)                         // tidak ditampilkan jika tidak ada

}
