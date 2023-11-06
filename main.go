package main

import "fmt"

func main() {
	// variable declartion
	
	var name string
	name = "hrishikesh"

	const gender string = "Male"
	var age int = 21
	loaction := "Thane"

	// if variable/pakage is not used, gives error
	fmt.Println(name, gender, age, loaction)
	
	// print statements
	fmt.Printf("Hello %v ", name)
	fmt.Printf("Age: %v", age)

	// taking input from user
	var surname string

	fmt.Scan(&surname)

	fmt.Printf("surname %v", surname)
	
}
