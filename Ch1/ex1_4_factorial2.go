package main

import "fmt"

func factorial2(n int) int {
	result := 1

	for n > 0 { // 반복문에 괄호가 없음
		result *= n
		n--
	}

	return result
}

func main() {
	fmt.Println(factorial2(5))
}
