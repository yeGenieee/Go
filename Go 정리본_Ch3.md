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



### 3.1.6 문자열을 숫자로

- strconv 패키지에 있는 함수를 이용
  - strconv.Atoi()* : 문자열을 정수로 바꿔줌
  - strconv.ParseInt() : 64비트 정수, 10진수가 아닌 수를 변환 가능
  - strconv.Itoa() / strconv.FormatInt() : 숫자열을 문자열로 변환



## 3.2 배열과 슬라이스

### 3.2.1 배열

- 연속된 메모리 공간을 순차적으로 이용하는 자료구조

- 크기가 자료형에 고정되어 있음

  ```go
  var str [3]string
  ```



### 3.2.2 슬라이스

- 배열보다 유연한 구조

- 길이와 용량을 갖고 있고, 길이가 변할 수 있는 구조

  ```go
  var fruits []string
  ```

  - 빈 슬라이스에는 nil 값이 들어감
  - C++의 nullptr, Java의 null과 비슷



  - 빈 스트링을 n개 갖고 있는 슬라이스

    ```go
    fruits := make([]string, n)
    ```

    - make를 통해 크기를 잡고 만든 슬라이스
      - 해당 자료형의 기본값이 들어감
        - 문자형 : "" 이 들어감
        - 정수형 : 0 이 들어감

- Go에서는, Python과 다르게, 슬라이싱에서 음수를 쓸 수 없음

- 슬라이싱 범위를 넘어가는 경우

  - Panic이 발생함 (다른 언어들의 exception과 비슷)
  - 에러를 반환할 수 있는 경우에는 패닉을 활용하지 않고 에러를 반환하는 방식을 이용함

### 3.2.3 슬라이스 덧붙이기

- std::vector, java.util.ArrayList와 비슷하게 Go에서도 슬라이스를 덧붙일 수 있음

  ```go
  fruits = append(fruits, "포도")
  ```

  - **append**를 통해 덧붙임



- 두 슬라이스 이어 붙이기

  ```go
  f1 := []string{"사과", "바나나", "토마토"}
  f2 := []string{"포도", "딸기"}
  f3 := append(f1, f2...) // 이어붙이기
  ```

- append 함수

  - 가변 인자를 받는 함수

  > f1에 "포도"와 "딸기"를 추가하여 f3에 할당하려면, ```f3 := append(f1, "포도", "딸기")```

  - 만약, ```f3 := append(f1,f2)```

    > f2는 문자열이 아닌 문자열 슬라이스이기 때문에 f1 뒤에 추가할 수 없음
    >
    > 따라서, f2 뒤에 ```...```을 붙여 f2에 있는 각각의 값들을 인자를 늘어놓듯이 늘어 놓을 수 있게 됨


### 3.2.4 슬라이스 용량

- 슬라이스 : 연속된 메모리 공간을 활용하는 자료구조

  - 용량에 제한이 있음



  - ```go
    make([]int, 5)
    []int{0,0,0,0,0}
    ```

  - 위와 같이 공간을 미리 할당하거나 이미 초기화를 한 경우,

    - 길이 뿐만 아니라 용량 까지 5로 맞춰짐
    - 만약, 용량이 가득찬 상태에서 슬라이스 덧붙이기를 진행한다면, 용량이 부족하므로 슬라이스 전체를 복사하게 됨

  - 용량 확인 함수

    - ```cap(x)```

- 용량 : 뒤에 얼마나 덧붙일 공간이 있느냐에 따라 용량이 결정됨

- 용량을 미리 지정해서 슬라이스 생성하기

  ```go
  nums := make([]int, 3, 5)
  ```

  > length는 3, capacity는 5
  >
  > 미리 공간을 예약해두는 개념, N개까지 길이가 늘어나더라도 복사가 일어나지 않게 하고 싶은 경우에 이용
  >
  > C++ 벡터에서 reserve 사용하는 개념과 동일

  ```go
  nums := make([]int, 0, 15)
  ```

  > 15개 까지 미리 예약
  >
  > ```nums = append(nums, x)``` 해도 15개를 넘어가지 않으면 슬라이스 복사가 일어나지 않음



### 3.2.5 슬라이스의 내부 구현



- 슬라이스
  - 배열을 가리키고 있는 구조체
    - 시작 주소
    - 길이
    - 용량

- nums 슬라이스 뒤에 10을 추가하는 경우

```go
nums = append(nums, 10)
```

1. nums의 늘어난 길이가 용량을 초과하는지 검사
   1. len(nums) + 1 <= cap(nums)
      - 시작 위치에서 길이만큼 오른쪽으로 이동(새로운 값을 넣을 위치), 새로운 값을 집어넣고, 길이가 증가한 슬라이스를 반환함
   2. len(nums) + 1 > cap(nums)
      - 용량을 초과하는 경우
        - 더 큰 크기의 배열을 새로 만들고, 슬라이스도 거기에 맞게 고쳐 반환
2. 반환된 슬라이스를 nums에 다시 할당



### 3.2.6 슬라이스 복사

- 반복문을 통한 슬라이스 복사

```go
func Example_sliceCopy() {
	src := []int{30, 20, 50, 10, 40}
	dest := make([]int, len(src))

	for i := range src {
		dest[i] = src[i] // 슬라이스를 range로 반복 시, i에는 인덱스가 들어감
	}
}
```

- copy 함수를 통한 슬라이스 복사

```go
func Example_sliceCopy2() {
    src := []int{30, 20, 50, 10, 40}
    
    copy(dest, src)
}
```



- copy 함수는 len(src)와 len(dest) 를 비교하여 더 작은 값 만큼만 복사함
- copy를 통해 슬라이스 전체를 복사하려면 아래와 같은 코드가 필요함

```go
dest := make([]int, len(src))
```



- append 함수를 이용한 슬라이스 복사

```go
func Example_sliceCopy3() {
    src := []int{10,20,30,40,50}
    dest := append([]int(nil), src...)
}
```



- 배열 포인터의 복사

```go
func Example_sliceCopy5() {
    src := []int{10,20,30,40,50}
    dest := src
}
```

- 슬라이스의 원소들이 복사되는 것이 아님
- 배열 포인터, 길이, 용량이 복사되는 것
- src 원소들의 값을 변경하면 dest 원소들의 값도 변경됨



### 3.2.7 슬라이스 삽입 및 삭제

- Go 언어 에서는 삽입 삭제 메서드를 따로 제공하지 않음



### 3.2.8 스택

