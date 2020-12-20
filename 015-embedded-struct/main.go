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

	fmt.Println(group.Name, group.Admin.firstName, "Group members : ", group.Users[0].firstName, " and ", group.Users[1].firstName, ", is available: ", group.IsAvailable)

}

func displayGroup(user Group) {

}
