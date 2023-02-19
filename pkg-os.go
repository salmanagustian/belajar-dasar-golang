package main

import (
	"fmt"
	"os"
)

func main() {
	args := os.Args
	fmt.Println("Arguments :")
	fmt.Println(args)

	name, err := os.Hostname()
	if err == nil {
		fmt.Println("Hostname Gopher", name)
	} else {
		fmt.Println("Error", err.Error())
	}

	os.Setenv("APP_USERNAME", "Gopher")
	username := os.Getenv("APP_USERNAME")
	fmt.Println("Current Env :", username)
}
