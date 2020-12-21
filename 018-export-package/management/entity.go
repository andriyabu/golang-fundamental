package management

import "fmt"

//User is a struct
type User struct {
	ID        int
	FirstName string
	LastName  string
	Email     string
	IsActive  bool
}

// Group is a struct
type Group struct {
	Name        string
	Admin       User
	Users       []User
	IsAvailable bool
}

func (group Group) DisplayGroup() {
	fmt.Printf("Group Name: %s", group.Name)
	fmt.Println("")
	fmt.Printf("Group Admin: %s", group.Admin.FirstName)
	fmt.Println("")
	fmt.Printf("Group members : %d", len(group.Users))
	fmt.Println("")
	for _, user := range group.Users {
		fmt.Println(user.FirstName)
	}
	fmt.Printf("Is Available : %t", group.IsAvailable)
	fmt.Println("")
}
