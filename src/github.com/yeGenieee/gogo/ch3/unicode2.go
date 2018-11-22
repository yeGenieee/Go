package main

import "fmt"

func main(){
	for _, r := range("가갛힣") {
		fmt.Println(string(r),r)
		// 44059 - 44032 : 27 , 가 ~ 갛 : 28자
	}
}