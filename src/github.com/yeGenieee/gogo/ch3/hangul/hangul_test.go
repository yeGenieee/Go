package hangul

import "fmt"

//유니코드 문자 단위로 처리
func ExampleHasConsonantSuffix() {
	fmt.Println(HasConsonantSuffix("Go 언어"));
	fmt.Println(HasConsonantSuffix("그럼"));
	fmt.Println(HasConsonantSuffix("우리 밥 먹고 합시다"));

	// Output:
	// .
}

// 바이트 단위로 처리
func Example_printBytes() {
	s := "가나다"
	for i := 0; i < len(s); i++ {
		fmt.Printf("%x:", s[i]) // 16진수 숫자 형식으로 출력
	}

	fmt.Println()

	// Output:
	
}

// 바이트 단위로 스트링을 16진수로 찍어보고 싶은 경우 -> 바로 Printf에서 %x 이용하면 됨
func Example_printBytes2() {
	s := "가나다"
	fmt.Printf("%x\n",s)
	fmt.Printf("% x\n",s)
	// Output:
}

// 문자열을 아예 바이트 단위의 슬라이스로 변환 가능
func Example_modifyBytes() {
	b := []byte("가나다") // byte는 uint8의 alias(별칭) , 문자열 s를 byte 슬라이스로 형변환 가능
	b[2]++ // 문자열 바꿀 수 있음
	fmt.Println(string(b)) // 슬라이스에서 다시 문자열로 변환
	// Output:
}

// 3.1.5 문자열 잇기
func Example_strCat() {
    s := "abc"
    ps := &s
    s += "def"
    
    fmt.Println(s)
    fmt.Println(*ps)
    // Output:
}