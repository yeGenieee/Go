package main

import (
	"fmt"
	"strconv"
)

func sing(n int) {
	money := 10

	for i := 1; i < n; i++ {
		fmt.Println("타잔이 " + strconv.Itoa(money) + "원짜리 팬티를 입고, " + strconv.Itoa(money+10) + "원 칼을 차고 노래를 한다.")
		money += 10
	}
}

func main() {
	sing(5)
}
