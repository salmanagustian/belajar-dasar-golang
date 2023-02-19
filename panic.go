package main

import "fmt"

func endApp() {
	message := recover()

	if message != nil {
		fmt.Println("Error Application with message", message)
	}

	fmt.Println("End Application")
}

func runApp(error bool) {
	defer endApp()

	if error {
		panic("ERROR APP!") // <- ke throw pada baris ini
	}

	fmt.Println("Application Running...")
}

func main() {
	runApp(true)
	fmt.Println("salman")
}
