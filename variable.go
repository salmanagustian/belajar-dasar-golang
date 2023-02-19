package main

import "fmt"

func main() {
	// single variable
	var name string
	var nilai uint16

	name = "salman"
	fmt.Println(name)

	nilai = 200
	fmt.Println(nilai)

	var friendName = "budi"
	fmt.Println(friendName)

	// type inference hanya untuk deklarasi awal
	age := 30
	fmt.Println(age)

	// multiple variable
	var (
		firstName = "salman"
		lastName  = "agustian"
	)

	fmt.Println(firstName)
	fmt.Println(lastName)

	// variable constant
	// constant tidak bisa menggunakan type inference
	const nameConstant string = "m salman"
	const hobby = "playing games"

	const (
		a = "m salman"
		b = "playing games"
	)
}
