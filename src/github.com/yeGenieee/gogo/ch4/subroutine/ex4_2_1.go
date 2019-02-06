package main

import (
	"bufio"
	"fmt"
	"io"
	"strings"
)

/**
 * 4.2.1
 * 인자와 반환값이 없는 함수
 */
func printHello() {
	fmt.Println("Hello!")
}

/**
 * 4.2.1
 * 함수 리터럴만을 남긴 후 호출
 */
func Example_funcLiteral() {
	func() {
		fmt.Println("Hello!")
	}()
	// Output:
	// Hello!
}

/**
 * 4.2.1
 * 함수 리터럴을 printHello 변수에 담은 후 호출
 */
func Example_funcLiteralVar() {
	printHello := func() {
		fmt.Println("Hello!")
	}

	printHello()
	// Output
	// Hello!
}

/**
 * 4.2.2
 * 고계 함수 (함수를 넘기고 받는 함수)
 */
func ReadFrom(r io.Reader, f func(line string)) error {
	scanner := bufio.NewScanner(r)

	for scanner.Scan() {
		f(scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		return err
	}
	return nil
}

/**
 * 4.2.2
 * 함수 리터럴을 넣어 호출
 */
func ExampleReadFrom_Print() {
	r := strings.NewReader("bill\ntom\njane\n")
	err := ReadFrom(r, func(line string) {
		fmt.Println("(", line, ")")
	})

	if err != nil {
		fmt.Println(err)
	}
	// Output:
	// (bill)
	// (tom)
	// (jane)
}

/**
 * 4.2.3
 * 클로저 (외부에서 선언한 변수를 함수 리터럴 내에서 마음대로 접근할 수 있는 코드)
 */
func ExampleReadFrom_append() {
	r := strings.NewReader("bill\ntom\njane")
	var lines []string
	// ReadFrom에 넘기는 함수는 그 함수가 이용하는 변수도 함께 가지고 넘어가게 됨
	err := ReadFrom(r, func(line string) {
		lines = append(lines, line) // ReadFrom에 넘겨주는 함수 리터럴 안에서 lines 라는 문자열 슬라이스를 사용중
	})

	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(lines)

	// Output:
	// [bill tom jane]
}

/**
 * 4.2.4
 * Generator
 */
func NewIntGenerator() func() int {
	var next int
	return func() int {
		next++
		return next
	}
}

/**
 * 4.2.4
 * 함수를 반환하는 고계 함수 NewIntGenerator
 */
func ExampleNewIntGenerator() {
	gen := NewIntGenerator()
	fmt.Println(gen(), gen(), gen(), gen(), gen())
	fmt.Println(gen(), gen(), gen(), gen(), gen())
	// Output:
	// 1 2 3 4 5
	// 6 7 8 9 10
}

/**
 * 4.2.4
 * gen1, gen2에 묶여있는 next
 */
func ExampleNewIntGenerator_multiple() {
	gen1 := NewIntGenerator()
	gen2 := NewIntGenerator()
	fmt.Println(gen1(), gen1(), gen1())
	fmt.Println(gen2(), gen2(), gen2(), gen2(), gen2())
	fmt.Println(gen1(), gen1(), gen1(), gen1())
	// Output:
	// 1 2 3
	// 1 2 3 4 5
	// 4 5 6 7
}
