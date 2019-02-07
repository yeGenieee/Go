# Ch1 시작하기

## 1.1 Go 언어 소개

- Go는 범용 프로그래밍 언어
- 깔끔하고 간결하게 생산성 높은 프로그래밍을 만들 수 있음
- 작성한 코드를 빠르게 컴파일하고 가비지 컬렉션을 지원
- 정적 자료형 언어이지만 동적 자료형 언어로 프로그램을 작성하는 듯한 느낌을 줌
- 생산성이 높은 이유?
  - 부분적이지만, 편리한 자료형 추론을 제공하여 굳이 반복해서 자료형 이름을 쓰지 않아도 됨
  - 소스 코드 형식을 자동으로 맞춰주는 도구 및 여러 편리한 도구가 기본으로 제공됨
  - Example 테스트를 이용하여 쉽게 테스트 코드를 작성하면서 코드 문서화까지 할 수 있음
  - 함수 리터럴 및 클로저를 자유자재로 사용할 수 있음
  - 명시적으로 인터페이스를 지정하지 않아도 인터페이스 구현이 가능
    - 기존에 있던 코드를 고치지 않고도 유연한 구현이 가능
  - 채널을 이용하여 동시성 구현을 락을 이용하지 않고 간편하게 할 수 있음
  - 언어 고유의 지원으로 교착 상태나 경쟁 상태를 파악하기 쉬움
  - 컴파일 속도가 빨라 컴파일 및 테스트를 반복적으로 수행하며 코드를 작성하기 용이
  - 가비지 컬렉션 지원으로 메모리 관리에 대한 부담을 덜 수 있음
  - 자료형 리터럴을 쉽게 쓸 수 있음



## 1.2 Hello World

```go
package main // 패키지 이름

import "fmt" // 가져올 모듈의 경로

// 메인 함수 (main 패키지에 있는 main 함수가 프로그램의 시작점)
func main() {
    fmt.Println("Hello World!")
}
```

- Go 컴파일러는 세미콜론으로 구분된 코드를 해석
- 구문 분석기가 소스 코드를 스캔하는 과정에서 규칙을 적용하여 자동으로 세미콜론을 붙이게 됨
- 줄이 끝난 것 처럼 보이기만 해도 세미콜론이 붙어서 컴파일러에게 전달됨



- 유의 사항

  ```go
  func main()
  {
      fmt.Println("Hello World")
  }
  ```

  - 2행이 닫는 괄호로 끝나기 때문에 단순한 전처리기는 세미콜론을 이 뒤에 붙여 구문 오류가 발생

- 주석
  - //
  - /**/



## 1.3 자료형 및 변수

- Go 언어는 자료형을 정적으로 검사하기 때문에 자료형이 정해져있음
- 정적 자료형을 지원하는 언어에서는 변수의 자료형을 미리 선언해주고 해당 변수에 맞는 형태의 값을 담는 방법을 이용
- Go 언어에는 정적 자료형을 지원하지만 자료형을 미리 선언하고 값을 할당하는 번거로움을 벗어나게 해주는 **자료형 추론 기능** 이 존재
- 정적 자료형을 지원하는 것은 변함없으므로 해당 변수에 다른 자료형의 값을 담을 수는 없음
- 그러나, 자료형의 이름을 써주지 않아도 됨



### 1.3.1 변수 선언

- 정수 x를 선언

  - 기존 C, C++, Java

    ```java
    int x;
    ```

  - Go

    ```go
    var x int
    ```

    

  - 베이직과 스칼라에서의 변수 선언 방식

  

  - 베이직과 SQL 등

    ```basic
    Dim x As Integer
    DECLARE @X AS INT
    ```

  - 스칼라, 파스칼, 에이다

    ```scala
    var x :int;
    var x : integer;
    ```



- 배열의 선언

  - C, C++

    ```C++
    int arr[5];
    ```

  - Java

    ```java
    int[] arr = new int[5];
    ```

  - Go

    ```go
    var arr [5]int
    ```

- Go 에서의 복잡한 함수 형태의 표현

  ```go
  func(int, int) // 두 정수를 인자로 받는 함수
  func(int) int // 하나의 정수를 인자로 받고, 하나의 정수를 반환하는 함수
  func(int, func(int, int)) func(int) int // 정수와 두 정수를 받아들이는 함수를 받고, 정수 하나를 받고 정수 하나를 반환하는 함수를 반환하는 함수
  ```



- 포인터

  - 정수 포인터

    ```go
    var p *int
    ```



- 변수 선언과 동시에 값을 할당

  ```go
  var x int = 10
  ```



### 1.3.2 자료형 추론

- Go 언어에서는 자료형이 무엇인지 알 수 있는 경우에는 자료형을 쓰지 않아도 됨

  ```go
  var i = 10
  var p = &i // i의 주소값을 p에 저장
  ```

- var 의 생략 (`=`대신에 `:=` 을 써야함)

  ```go
  i := 10
  p := &i
  ```

- C++ (정적 자료형)

  ```c++
  auto i = 10; // 컴파일 시점에 i가 정수형임을 앎 (int i=10;)
  auto* p = &i;
  ```

- Python (동적 자료형)

  ```python
  i = 10 # i에 할당되는 값에 따라 자료형이 정해짐
  ```



- Go 언어는 변수에 정적으로 자료형이 고정됨

  ```go
  i := 10 // 가능, i는 새로운 정수형 변수
  s := "hello" // 가능, s는 새로운 문자열 변수
  i = 20 // 가능, i의 값을 변경
  j = 30 // 불가능, j는 선언되지 않은 변수
  i = "hello" // 불가능, 이미 i는 정수형 변수임
  i := 30 // 불가능, 이미 i는 선언되었음
  i := "hi" // 불가능, 마찬가지 이유
  ```

  ```go
  package main
  
  import "fmt"
  
  func main() {
      var a = 10
      b := "little"
      b += " Gophers"
      fmt.Println("Hello, playground", a, b)
  }
  ```



## 1.4 함수와 간단한 제어 구조

- 팩토리얼

  ```go
  package main
  
  import "fmt"
  
  // fac returns n!
  func fac(n int) int {
      if n <= 0 {
          return 1
      } 
      return n * fac(n-1)
  }
  
  func main() {
      fmt.Println(fac(5))
  }
  ```

  - 함수의 반환 자료형은 함수 뒤에 써줌

- 반복문을 이용한 팩토리얼

  ```go
  package main
  
  import "fmt"
  
  // facItr returns n!
  func facItr(n int) int {
      result := 1
      for n > 0 {
          result *= n
          n--
      }
      return result 
  }
  
  func facItr2(n int) int {
      result := 1
      for i := 2; i <=n; i++ {
          result *= i
      }
      return result
  }
  
  func main() {
      fmt.Println(facItr(5))
  }
  ```

  







