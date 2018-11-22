package main

import "fmt"

func main(){
	for i,r := range "가나다" {
		fmt.Println(i, r) // 각각의 유니코드 문자에 대해 바이트 위치를 i에 넣고, 각 문자를 r에 넣음
		// i가 0, 3, 6인 이유 : UTF-8로 "가나다" 표현하는 데 글자당 3바이트가 필요하기 때문
		// i는 int, r은 rune형 (int32의 별칭-유니코드 포인트 하나를 담을 수 있음)
		fmt.Println(string(r),r) // r을 해당 코드 포인트의 유니코드 문자열로 바꿈
	}

	fmt.Println(len("가나다")) // len은 각 글자당 3바이트 이기 때문에, 총 9 바이트가 됨
}