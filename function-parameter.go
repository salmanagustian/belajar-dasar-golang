package main

import "fmt"

type AgeFilter func(int) string

func sayHelloWithFilter(name string, age int, filter AgeFilter) {
	ageFiltered := name + " " + filter(age)

	fmt.Println("Hey " + ageFiltered)
}

func ageFilter(age int) string {
	if age < 18 {
		return "you are age restricted"
	} else {
		return "you are allowed"
	}
}

func main() {
	sayHelloWithFilter("Salman", 23, ageFilter)
}
