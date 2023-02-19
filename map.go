package main

import "fmt"

func main() {
	person := map[string]string{
		"name":    "salman",
		"address": "Cimahi",
		"phone":   "628212297823",
	}

	fmt.Println(person)
	fmt.Println(person["name"])
	fmt.Println(person["address"])

	delete(person, "phone")

	fmt.Println(person)

	var book map[string]string = make(map[string]string)
	book["title"] = "Ikigai"
	book["author"] = "John"

	fmt.Println(book)
	fmt.Println(len(book))

	var person2 map[string]int = make(map[string]int)
	person2["age"] = 30
	person2["height"] = 170
	person2["weight"] = 65

	fmt.Println(person2)
	fmt.Println("person 2 weight is :", person2["weight"])
}
