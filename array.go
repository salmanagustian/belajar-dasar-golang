package main

import "fmt"

func main() {

	var friendNames [3]string

	friendNames[0] = "Ahmad"
	friendNames[1] = "Pran"
	friendNames[2] = "Yogi"

	fmt.Println(friendNames)

	var hobbies = [3]string{
		"Games",
		"Chess",
		"Listen Music",
	}

	fmt.Println(hobbies)

}
