package main

import "fmt"

func main() {
	sayGoodBye := getGoodBye

	result := sayGoodBye("Salman")
	fmt.Println(result)
	fmt.Println(sayGoodBye("Agustian"))
}

func getGoodBye(name string) string {
	return "Good Bye " + name
}
