package main

import "fmt"

func main() {
	result := sayHello()
	fmt.Println(result)

	var multiple = multiple(5, 10)
	fmt.Println(multiple)

	// function return multiple values
	firstName, lastName := getFullName()
	fmt.Println(firstName)
	fmt.Println(lastName)

	// function named return values
	firstName2, lastName2 := getCompleteName()
	fmt.Println(firstName2)
	fmt.Println(lastName2)
}

func sayHello() string {
	return "Hello World"
}

func multiple(a int, b int) int {
	return a * b
}

func getFullName() (string, string) {
	return "Muhammad", "Salman"
}

func getCompleteName() (firstName string, lastName string) {
	firstName = "Ahmad"
	lastName = "Hasyim"

	return
}
