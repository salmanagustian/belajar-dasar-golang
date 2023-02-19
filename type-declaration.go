package main

import "fmt"

func main() {
	type NoKTP string
	type Married bool
	var identityNumber NoKTP = "32910020000"
	var marriedStatus Married = true

	fmt.Println(identityNumber)
	fmt.Println(marriedStatus)
}
