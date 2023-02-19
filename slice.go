package main

import (
	"fmt"
)

func main() {

	var months = [...]string{
		"januari",
		"februari",
		"maret",
		"april",
		"mei",
		"juni",
		"juli",
		"agustus",
		"september",
		"oktober",
		"november",
		"desember",
	}

	var slice = months[11:]
	fmt.Println(slice)
	fmt.Println(len(slice))
	fmt.Println(cap(slice))
	fmt.Println(months)

	var slice2 = append(slice, "item 1", "item 2")
	fmt.Println("====")
	// slice2[0] = "Bukan Maret"
	fmt.Println(slice2)
	fmt.Println(months)

	newSlice := make([]string, 2, 5)
	newSlice[0] = "item 1"
	newSlice[1] = "item 2"
	fmt.Println(newSlice)

	copySlice := make([]string, len(newSlice), cap(newSlice))

	copy(copySlice, newSlice)
	fmt.Println(copySlice)

	iniArray := [5]int{1, 2, 3, 4, 5}
	iniSlice := []int{1, 2, 3, 4, 5}

	fmt.Println(iniArray)
	fmt.Println(iniSlice)
}
