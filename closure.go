package main

import "fmt"

func main() {
	name := "salman"
	counter := 0

	increment := func() {
		name := "agustian"
		fmt.Println("Increment")
		counter++
		fmt.Println(name)
	}

	increment()
	increment()

	fmt.Println(counter)
	fmt.Println(name)
}
