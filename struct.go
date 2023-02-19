package main

import "fmt"

type Customer struct {
	Name, Address string
	Age           int
}

func getCustomer(data Customer) string {
	return "Customer with name " + data.Name
}

func main() {
	var data Customer
	data.Name = "M Salman Agustian"
	data.Address = "Cimahi"
	data.Age = 26

	fmt.Println(data)

	data2 := Customer{
		Name:    "Salman",
		Address: "Indonesia",
		Age:     23,
	}

	fmt.Println(data2)

	customer := getCustomer(data)
	fmt.Println(customer)
}
