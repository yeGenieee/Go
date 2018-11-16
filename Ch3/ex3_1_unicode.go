package main

import "fmt"

func main() {
	for i, r := range "가나다" {
		// 각 유니코드 문자에 대해 바이트 위치에 i를 넣고, 각 문자를 r에 넣음 (i가 0,3,6인 이유는 글자당 3바이트가 필요하기 때문)
		// i는 int형, r은 rune형  (int32)
		fmt.Println(i, r)
		fmt.Println(string(r), r) // r을 해당 코드 포인트의 유니코드 문자열로 변환
	}

	fmt.Println(len("가나다")) // 문자열의 길이
}
