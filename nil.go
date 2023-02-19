package main

import "fmt"

func newMap(name string) map[string]string {
	if name == "" {
		return nil
	} else {
		return map[string]string{
			"name": name,
		}
	}
}

func main() {
	// nil dapat digunakan untuk tipe data seperti interface, function, map, slice, pointer dan slice
	var person map[string]string = newMap("Momo")

	// if person["name"] == ""

	if person == nil {
		fmt.Println("Data Kosong")
	} else {
		fmt.Println(person)
	}

}
