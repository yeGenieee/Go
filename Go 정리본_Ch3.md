# Ch3 문자열 및 자료구조

- 문자열 처리
- 배열과 슬라이스
- 맵
- 기본적인 입출력



## 3.1 문자열

- 바이트들이 연속적으로 나열되어 있는 것
- string 이용
  - string은 읽기 전용
  - 바이트 조작이 불가능
- []byte
  - 바이트 배열

### 3.1.1 유니코드 처리

- Go 언어 : UTF-8 인코딩

  > **UTF-8**
  >
  > 인코딩의 한 방식
  >
  > 유니코드 포인트를 나타내기 위한 바이트 수가 가변적
  >
  > > 0 ~127까지는 1 Byte로 표현됨
  > >
  > > 그 이후부터는 2, 3, ... 6 Byte 까지 크기가 늘어날 수 있음



```go
package main

import "fmt"

func main(){
	for i,r := range "가나다" {
		fmt.Println(i, r)
	}

	fmt.Println(len("가나다"))
}
```

- for i := range
  - for each, for ... in ... 같은 형태의 반복문



#### Go test 해보기

- {파일명}_test.go 작성

  - **_test** 를 붙여야 함

- 테스트 함수는 반드시 **Example**로 시작해야 함

  ```go
  go test github.com/yeGenieee/gogo/ch3/hangul
  ```

  - go test
    - go run 과 유사하지만 테스트를 수행함



### 슬라이스

- 배열을 조금 더 유연하게 만든 구조
- C++의 std::vector,  자바의 java.util.ArrayList 와 유사



### 3.1.4 패키지 문서

##### https://golang.org/pkg

Go에서 지원하는 모든 표준 라이브러리가 문서화 되어 있음



### 3.1.5 문자열 잇기

- Go의 string : 읽기 전용

- **+** 연산으로 문자열 concatenation

  ```go
  func Example_strCat() {
      s := "abc"
      ps := &s
      s += "def" // s 뿐만 아니라 ps가 가리키고 있는 값까지 변경됨
      
      fmt.Println(s)
      fmt.Println(*ps)
      // Output:
  }
  ```

  - ps : s의 주소값을 취한 포인터 형식인 *string 형
  - s += "def"
    - 간단한 문자열 이어 붙이는 용도로 좋음

- fmt 패키지의 Sprint 를 이용하여도 concat 가능

  ```go
  s = fmt.Sprint(s, "def")
  s = fmt.Sprintf("%sdef", s)
  ```

  - Sprint 
    - 문자열이 아닌 다른 것들도 이어 붙일 수 있음
  - Sprintf
    - 형식 지정 가능

- strings 패키지 이용도 가능

  ```go
  s = strings.Join([]string {s, "def"}, " ")
  ```

