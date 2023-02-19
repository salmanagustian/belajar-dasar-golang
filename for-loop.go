package main

import "fmt"

func main() {
	// for loops
	counter := 1

	for counter <= 10 {
		fmt.Println("Perulangan ke ", counter)
		counter++
	}

	// for loops as while
	saldo := 10000
	harga := 3500
	counterPembelian := 0

	for saldo >= harga {
		counterPembelian++
		saldo -= harga
		fmt.Println("Pembelian yang ke ", counterPembelian, "dengan sisa saldo adalah ", saldo)
	}

	// for statement
	for counter := 1; counter <= 10; counter++ {
		fmt.Println("Perulangan ke ", counter)
	}

	// for slice
	slice := []string{
		"Salman", "Ahmad", "Febri",
	}

	for i := 0; i < len(slice); i++ {
		fmt.Println(slice[i])
	}

	// for range biasanya digunakan untuk melakukan iterasi semua data collection
	// contoh nya array, slice, map

	names := []string{"momo", "oreo", "genero"}

	for _, val := range names {
		fmt.Println(val)
	}

	person := map[string]string{
		"name":    "salman",
		"address": "Cimahi",
		"phone":   "628212297823",
	}

	for key, val := range person {
		fmt.Println(key, "=", val)
	}
}
