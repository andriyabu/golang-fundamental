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

// Group is a struct
type Group struct {
	Name        string
	Admin       User
	Users       []User
	IsAvailable bool
}

func (group Group) displayGroup() {
	fmt.Printf("Group Name: %s", group.Name)
	fmt.Println("")
	fmt.Printf("Group Admin: %s", group.Admin.firstName)
	fmt.Println("")
	fmt.Printf("Group members : %d", len(group.Users))
	fmt.Println("")
	for _, user := range group.Users {
		fmt.Println(user.firstName)
	}
	fmt.Printf("Is Available : %t", group.IsAvailable)
	fmt.Println("")
}

func main() {

	user := User{1, "andri", "yabu", "andriyabu@gmail.com", true}
	user2 := User{
		firstName: "Abeleon",
		lastName:  "Godwin Pekambani",
		email:     "abeleongodwin@gmail.com",
		isActive:  true,
	}

	users := []User{user, user2}
	group := Group{"Gamer", user, users, true}

	group.displayGroup()

}
