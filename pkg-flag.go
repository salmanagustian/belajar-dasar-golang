package main

import (
	"flag"
	"fmt"
)

func main() {
	host := flag.String("host", "localhost", "Put your Host")
	username := flag.String("username", "root", "Put your Username")
	password := flag.String("password", "secret", "Put your Password")
	number := flag.Int("number", 100, "Put your Number")

	flag.Parse()

	fmt.Println("Host :", *host)
	fmt.Println("Username :", *username)
	fmt.Println("Password :", *password)
	fmt.Println("Number :", *number)
}
