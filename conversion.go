package main

import "fmt"

func main() {
	var nilai32 int32 = 32768
	var nilai64 = int64(nilai32)
	var nilai16 = int16(nilai64)

	fmt.Println(nilai16)

	var name = "salman"
	var e byte = name[1]
	var eString string = string(e)

	fmt.Println(e)
	fmt.Println(eString)
}
