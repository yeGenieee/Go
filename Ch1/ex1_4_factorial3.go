package main

import "fmt"

func factorial3(n int) int {
	result := 1

	for i := 2; i <= n; i++ {
		result *= i
	}

	return result
}

func main() {
	fmt.Println(factorial3(5))
}
