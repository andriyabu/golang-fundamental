package main

import (
	"export-package/management"
)

func main() {

	user := management.User{
		ID:        1,
		FirstName: "andri",
		LastName:  "yabu",
		Email:     "andriyabu@gmail.com",
		IsActive:  true,
	}
	user2 := management.User{
		ID:        2,
		FirstName: "Abeleon",
		LastName:  "Godwin Pekambani",
		Email:     "abeleongodwin@gmail.com",
		IsActive:  true,
	}

	users := []management.User{
		user,
		user2,
	}
	group := management.Group{
		Name:        "Gamer",
		Admin:       user,
		Users:       users,
		IsAvailable: true,
	}

	group.DisplayGroup()

}
