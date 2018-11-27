package main

import (
	"strconv"

	"fmt"
)

func main(){
	var i int
	var k int64
	var f float64
	var s string
	var err error

	i, err = strconv.Atoi("350") // i == 350
	k, err = strconv.ParseInt("cc7fdd", 16, 32) // k == 13402077
	k, err = strconv.ParseInt("0xcc7fdd", 0, 32) // k == 13402077
	f, err = strconv.ParseFloat("3.14", 64) // f == 3.14
	s = strconv.Itoa(340)
	s = strconv.FormatInt(13402077, 16) // s == "cc7fdd"

	var num int
	fmt.Sscanf("57", "%d", &num) // num == 57, sscanf가 정수형 변수 num에 값을 변경해야 하므로 포인터를 넘겨주어야 함

	var str string
	str = fmt.Sprint(3.14) // s == "3.14"
	str = fmt.Sprintf("%x", 13402077) // s == "cc7fdd"

}