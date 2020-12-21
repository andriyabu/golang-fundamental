package main

import (
	"fmt"
)
//Student is a struct
type Student struct{
	ID int
	Name string
	Gpa float32
}

func (student *Student) change()  {
	student.Name = student.Name + " S.Kom"
}

func main()  {

	student:= Student{1,"Andri yabu",4.0}

	student.change()

	fmt.Println(student.Name)
	
}