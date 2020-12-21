package main

import (
	"export-package/management"
)

func main() {

	user := management.User{1, "andri", "yabu", "andriyabu@gmail.com", true}
	user2 := management.User{
		FirstName: "Abeleon",
		LastName:  "Godwin Pekambani",
		Email:     "abeleongodwin@gmail.com",
		IsActive:  true,
	}

	users := []management.User{user, user2}
	group := management.Group{"Gamer", user, users, true}

	group.DisplayGroup()

}
