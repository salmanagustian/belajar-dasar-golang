package main

import "fmt"

func logging() {
	fmt.Println("End Application")
}

func runApplication(value int) {
	defer logging()
	fmt.Println("Start Application")

	result := 10 / value

	fmt.Println("Result", result)
}

func main() {
	runApplication(10)
}
