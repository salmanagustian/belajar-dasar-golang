package main

import "fmt"

type Address struct {
	City, Province, Nationality string
}

func changeAddressToIndonesia(address *Address) {
	address.Nationality = "Indonesia"
}

func main() {
	/*
		Pointer
		terdapat dua jenis yaitu pass by value dan pass by reference
		pass by value artinya isi variable 1 di "copy" ke variable 2
	*/
	var address1 Address = Address{"Cimahi", "Jawa Barat", "Indonesia"}
	var address2 *Address = &address1
	var address3 *Address = &address1

	// membuat pointer dari variable baru
	// var address3 *Address = &Address{"Malang", "Jawa Timur", "Indonesia"}

	address2.City = "Bandung"

	// tanpa operator "*"
	// address2 = &Address{"Jakarta", "DKI Jakarta", "Indonesia"}

	// dengan operator "*" saat assigment object baru
	*address2 = Address{"Jakarta", "DKI Jakarta", "Indonesia"}

	// var address4 *Address = &address1

	// membuat pointer dari variable baru menggunakan keyword new
	// var address5 *Address = new(Address)
	// address5.City = "Subang"

	fmt.Println(address1)
	fmt.Println(address2)
	fmt.Println(address3)
	// fmt.Println(address4)
	// fmt.Println(address5)

	var alamat Address = Address{"Bandung", "Jawa Barat", ""}
	var alamatPointer *Address = &alamat
	changeAddressToIndonesia(alamatPointer)

	fmt.Println(alamat)
}
