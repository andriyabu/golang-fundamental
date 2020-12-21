package main

import (
	"fmt"
)

//Student is a struct
type Student struct {
	ID   int
	Name string
	Gpa  float32
}

func main() {
	student := Student{1, "andri yabu", 3.8}

	change(&student)
	fmt.Println(student.Name)
}

func change(student *Student) {
	student.Name = student.Name + " S.Kom"
}
