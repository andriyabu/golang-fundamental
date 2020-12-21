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

// method display
func (user User) display() string  {
	return fmt.Sprintf("Name : %s %s , Email: %s , is active: %t",user.firstName,user.lastName,user.email,user.isActive)

}


func main() {

	user := User{1, "andri", "yabu", "andriyabu@gmail.com", true}

	result := user.display()
	fmt.Println(result)

	user2 := User{
		firstName: "Abeleon",
		lastName:  "Godwin Pekambani",
		email:     "abeleongodwin@gmail.com",
		isActive:  true,
	}

	fmt.Println(user2.display())

}


