package main

import "fmt"

func factorialLoop(value int) int {
	result := 1

	for i := value; i > 0; i-- {
		result *= i
	}

	return result
}

func factorialRecursive(value int) int {
	if value == 1 {
		return 1
	} else {
		return value * factorialRecursive(value-1)
	}
}

func launchRocket(counter int) int {
	if counter == 1 {
		return 1
	} else {
		fmt.Println(counter)
		return launchRocket(counter - 1)
	}
}

func main() {
	loop := factorialLoop(5)
	fmt.Println(loop)

	fmt.Println(factorialRecursive(5))
	fmt.Println(launchRocket(5))
	// 5 * 4 * 3 * 2 * 1
}
