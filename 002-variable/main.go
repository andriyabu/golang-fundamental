package main

import (
	"fmt"
)

func main() {
	var name string // deklarasi variable bisa langsung diisi ex: var name string = "andri yabu"
	name = "Andri yabu"

	// tanpa menulis kata var dan tipe datanya, tanda := adalah inisialisasi awal
	umur := 21

	// mengubah isi variable hanya menggunakan = bukan :=
	umur = 35

	fmt.Println("nama saya ", name, " umur:", umur)
}
