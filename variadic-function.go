package main

import "fmt"

func main() {
	total := sumAll(10, 10, 10, 10, 5)
	fmt.Println(total)

	slice := []int{10, 10, 5}
	total = sumAll(slice...)
	fmt.Println(total)
}

// variadic function or varargs
func sumAll(numbers ...int) int {
	total := 0

	for _, number := range numbers {
		total += number
	}

	return total
}
