package main

import "fmt"

func factorial2(n int) int {
	result := 1

	for n > 0 {
		result *= n
		n--
	}

	return result
}

func main() {
	fmt.Println(factorial2(5))
}
