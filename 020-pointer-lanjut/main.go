package main

import (
	"fmt"
)


func main()  {

	var numberA int = 5
	// var numberB int = 10

	fmt.Println(numberA)
	fmt.Println("Memori address : ",&numberA)
	// fmt.Println(numberB)

	change(&numberA, 100)
	
	fmt.Println(numberA)

}

func change(old *int, new int)  {
	*old = new
	fmt.Println("nilai didalam function : ",*old)
	fmt.Println("memory addres: ",&old)
}