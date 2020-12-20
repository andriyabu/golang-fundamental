package main

import (
	"fmt"
)

//User is a struct
type User struct {
	ID        int
	firstName string
	lastName  string
	email     string
	isActive  bool
}

func main() {
	// cara pertama
	user := User{}
	user.ID = 1
	user.firstName = "andri"
	user.lastName = "yabu"
	user.email = "andriyabu@gmail.com"
	user.isActive = true

	// cara kedua bisa acak urutannya
	user2 := User{
		isActive:  true,
		ID:        2,
		lastName:  "godwin",
		firstName: "abeleon",
		email:     "abeleongodwin@gmail.com",
	}

	// cara ketiga langsung isi data tapi harus sesuai urutan
	user3 := User{1, "nadia", "hamu meha", "nadia@gmail.com", true}

	fmt.Println(user.firstName, user.lastName, user.email, user.isActive)
	fmt.Println(user2)
	fmt.Println(user3)
}
