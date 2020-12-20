package main

import (
	"fmt"
)

func main() {
	students := []map[string]string{
		{"name": "andri yabu", "gender": "male"},
		{"name": "abeleon", "gender": "male"},
		{"name": "nadia", "gender": "female"},
	}

	// fmt.Println(students)

	for _, student := range students {
		fmt.Println("Name:", student["name"], "\tGender: ", student["gender"])
	}
}
