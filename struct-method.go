package main

import "fmt"

type Customer struct {
	Name, Address string
	Age           int
}

type Customer2 struct {
	Age int
}

func (customer Customer) sayHi() {
	fmt.Println("Hello", customer.Name)
}

func (customer Customer) sayHey(name string) {
	fmt.Println("Hello", name, "My Name is", customer.Name)
}

func main() {
	cust := Customer{Name: "m salman"}
	cust.sayHi()

	cust.sayHey("Azis")
}
